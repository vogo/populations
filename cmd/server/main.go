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

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vogo/populations"
)

// API响应结构
type PopulationResponse struct {
	Status     string  `json:"status"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Radius     float64 `json:"radius_meters"`
	Population float64 `json:"population"`
	Message    string  `json:"message,omitempty"`
}

// 错误响应结构
type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// 全局处理器
var processor *populations.Processor

func main() {
	// 1. 初始化人口数据处理器
	tiffPath := os.Getenv("FUTURE_POP_TIFF_PATH")
	if tiffPath == "" {
		log.Fatal("FUTURE_POP_TIFF_PATH env not set")
	}

	var err error
	processor, err = populations.NewProcessor(tiffPath)
	if err != nil {
		log.Fatalf("init processor error: %v", err)
	}
	defer processor.Close()

	log.Println("future population data loaded, start api server...")

	// 2. 设置路由
	r := mux.NewRouter()

	// 健康检查端点
	r.HandleFunc("/health", healthHandler).Methods("GET")

	// 人口查询端点
	r.HandleFunc("/api/v1/population", populationHandler).Methods("GET")

	// 3. 启动服务器
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080"
	}
	log.Printf("serve at %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}

// 健康检查
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

// 人口查询处理
func populationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 解析查询参数
	query := r.URL.Query()

	latStr := query.Get("lat")
	lngStr := query.Get("lng")
	radiusStr := query.Get("radius")

	if latStr == "" || lngStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrorResponse{
			Status: "error",
			Error:  "missing latitude or longitude parameter (lat, lng)",
		})
		return
	}

	// 转换参数
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil || lat < -90 || lat > 90 {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrorResponse{
			Status: "error",
			Error:  "latitude parameter invalid, should be in range -90 to 90",
		})
		return
	}

	lng, err := strconv.ParseFloat(lngStr, 64)
	if err != nil || lng < -180 || lng > 180 {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrorResponse{
			Status: "error",
			Error:  "longitude parameter invalid, should be in range -180 to 180",
		})
		return
	}

	// 默认半径3公里
	radius := 3000.0
	if radiusStr != "" {
		radius, err = strconv.ParseFloat(radiusStr, 64)
		if err != nil || radius <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(ErrorResponse{
				Status: "error",
				Error:  "invalid radius parameter, should be positive",
			})
			return
		}
	}

	// 计算人口
	population, err := processor.GetPopulationInRadius(lng, lat, radius)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	// 返回成功响应
	response := PopulationResponse{
		Status:     "success",
		Latitude:   lat,
		Longitude:  lng,
		Radius:     radius,
		Population: population,
		Message:    fmt.Sprintf("successfully calculate population in radius %.0f meters", radius),
	}

	_ = json.NewEncoder(w).Encode(response)
}
