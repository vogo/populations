/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package populations

import (
	"fmt"
	"math"

	"github.com/lukeroth/gdal"
)

// Processor 人口数据处理器
type Processor struct {
	dataset   gdal.Dataset
	band      gdal.RasterBand
	transform [6]float64 // 地理变换参数
	width     int
	height    int
}

// NewProcessor 创建新的处理器
func NewProcessor(tiffPath string) (*Processor, error) {
	fmt.Println("Registering GDAL drivers...")
	// Register all drivers
	// gdal.AllRegister()

	fmt.Printf("Opening TIFF file: %s\n", tiffPath)
	dataset, err := gdal.Open(tiffPath, gdal.ReadOnly)
	if err != nil {
		return nil, fmt.Errorf("gdal open error: %v", err)
	}
	fmt.Println("TIFF file opened successfully")

	transform := dataset.GeoTransform()

	// Check if transform is valid
	// Some datasets might not have a transform set, returning default [0, 1, 0, 0, 0, 1]
	// But for GeoTIFF it should be there.

	width := dataset.RasterXSize()
	height := dataset.RasterYSize()

	band := dataset.RasterBand(1) // Usually band 1 for population

	return &Processor{
		dataset:   dataset,
		band:      band,
		transform: transform,
		width:     width,
		height:    height,
	}, nil
}

// Close 关闭数据集
func (p *Processor) Close() {
	p.dataset.Close()
}

// GetPopulationInRadius 计算指定半径内的人口
func (p *Processor) GetPopulationInRadius(lon, lat, radiusMeters float64) (float64, error) {
	// 1. 将中心点从经纬度转换为像素坐标
	px, py := p.geoToPixel(lon, lat)

	// 2. 计算半径对应的像素距离
	radiusPixels := p.metersToPixels(radiusMeters, lat)

	// 3. 定义搜索边界
	xMin := int(math.Floor(float64(px) - radiusPixels))
	xMax := int(math.Ceil(float64(px) + radiusPixels))
	yMin := int(math.Floor(float64(py) - radiusPixels))
	yMax := int(math.Ceil(float64(py) + radiusPixels))

	// Clamp to image bounds
	if xMin < 0 {
		xMin = 0
	}
	if yMin < 0 {
		yMin = 0
	}
	if xMax >= p.width {
		xMax = p.width - 1
	}
	if yMax >= p.height {
		yMax = p.height - 1
	}

	if xMin > xMax || yMin > yMax {
		return 0, nil // Area is outside of image
	}

	winWidth := xMax - xMin + 1
	winHeight := yMax - yMin + 1

	// Read data for this window
	buffer := make([]float32, winWidth*winHeight)

	// IO(rw Flag, xOff, yOff, xSize, ySize int, buffer interface{}, bufXSize, bufYSize, pixelSpace, lineSpace int) error
	err := p.band.IO(gdal.Read, xMin, yMin, winWidth, winHeight, buffer, winWidth, winHeight, 0, 0)
	if err != nil {
		return 0, fmt.Errorf("read raster error: %v", err)
	}

	totalPopulation := 0.0

	// 4. 遍历边界内的所有像素
	for y := 0; y < winHeight; y++ {
		for x := 0; x < winWidth; x++ {
			// Global pixel coordinates
			globalX := xMin + x
			globalY := yMin + y

			// 计算像素中心点的地理坐标
			pixelLon, pixelLat := p.pixelToGeo(float64(globalX)+0.5, float64(globalY)+0.5)

			// 检查是否在圆形半径内
			if p.distanceMeters(lon, lat, pixelLon, pixelLat) <= radiusMeters {
				// 获取像素值
				value := float64(buffer[y*winWidth+x])
				if value >= 0 { // 忽略NoData值（通常为负值）
					totalPopulation += value
				}
			}
		}
	}

	return totalPopulation, nil
}

// geoToPixel 地理坐标转像素坐标
func (p *Processor) geoToPixel(lon, lat float64) (float64, float64) {
	// GDAL Transform:
	// X_geo = GT[0] + X_pixel * GT[1] + Y_pixel * GT[2]
	// Y_geo = GT[3] + X_pixel * GT[4] + Y_pixel * GT[5]
	//
	// Assuming GT[2] and GT[4] are 0 (no rotation), which is standard for most GeoTIFFs.
	// X_pixel = (X_geo - GT[0]) / GT[1]
	// Y_pixel = (Y_geo - GT[3]) / GT[5]

	x := (lon - p.transform[0]) / p.transform[1]
	y := (lat - p.transform[3]) / p.transform[5]
	return x, y
}

// pixelToGeo 像素坐标转地理坐标
func (p *Processor) pixelToGeo(x, y float64) (float64, float64) {
	lon := p.transform[0] + x*p.transform[1] + y*p.transform[2]
	lat := p.transform[3] + x*p.transform[4] + y*p.transform[5]
	return lon, lat
}

// metersToPixels 米转换为像素
func (p *Processor) metersToPixels(meters, lat float64) float64 {
	// 简化的转换：1度纬度约111km，1度经度约111km*cos(lat)
	metersPerDegreeLat := 111320.0
	metersPerPixelLat := math.Abs(p.transform[5] * metersPerDegreeLat)
	return meters / metersPerPixelLat
}

// distanceMeters 计算两点间距离（米）- Haversine公式
func (p *Processor) distanceMeters(lon1, lat1, lon2, lat2 float64) float64 {
	const R = 6371000 // 地球半径（米）

	φ1 := lat1 * math.Pi / 180
	φ2 := lat2 * math.Pi / 180
	Δφ := (lat2 - lat1) * math.Pi / 180
	Δλ := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(Δφ/2)*math.Sin(Δφ/2) +
		math.Cos(φ1)*math.Cos(φ2)*math.Sin(Δλ/2)*math.Sin(Δλ/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}
