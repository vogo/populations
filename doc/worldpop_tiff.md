# WorldPop TIFF 文件格式详细介绍

WorldPop 发布的空间人口数据主要使用 **GeoTIFF** 格式。这是一种基于 TIFF (Tagged Image File Format) 的标准光栅图像格式，嵌入了地理空间元数据，使其能够被 GIS (地理信息系统) 软件直接读取和定位。

## 1. TIFF 文件基础

TIFF 是一种灵活的位图图像格式，广泛用于存储图像数据。在 GIS 领域，TIFF 被扩展为 GeoTIFF，允许在文件头中存储地理坐标系、投影信息和像素比例等元数据。

### 文件结构
- **IFD (Image File Directory)**: 包含描述图像数据的标签 (Tags)。
- **图像数据**: 实际的像素值矩阵。

## 2. WorldPop TIFF 文件的具体特征

当您下载 WorldPop 的 `.tif` 文件（例如 `ppp_2020_1km_Aggregated.tif`）时，它通常具有以下技术规格：

### 2.1 数据类型与像素值
- **像素值含义**：
    -   对于 **PPP (People Per Pixel)** 数据集：每个像素的值代表该网格单元内的**估计人数**。这是一个浮点数 (Floating Point)，例如 `12.54` 表示该网格内估计有 12.54 人。
    -   对于 **PD (Population Density)** 数据集：每个像素的值代表**每平方公里的人口密度**。
-   **数据类型 (Data Type)**: 通常为 `Float32` (32位浮点型)，以保留估计值的精度。
-   **NoData 值**: 文件中会定义一个特定的值（通常是 `-9999` 或 `NaN`）来表示水体、无人区或无数据区域。在分析时必须排除这些值。

### 2.2 空间分辨率 (Spatial Resolution)
WorldPop 提供多种分辨率的数据：
-   **3弧秒 (3 arc-seconds)**: 约 **100米** (在赤道处)。这是 WorldPop 的核心高分辨率产品。
-   **30弧秒 (30 arc-seconds)**: 约 **1公里** (在赤道处)。通常由 100米数据聚合而来，便于大尺度分析。

### 2.3 坐标参考系统 (CRS)
-   **地理坐标系**: 绝大多数 WorldPop 数据使用 **WGS84 (EPSG:4326)** 地理坐标系。
-   **单位**: 度 (Degrees)。
-   这意味着像素的宽度和高度是以经纬度为单位的，而不是米。在进行面积计算或距离分析时，通常需要将其投影到投影坐标系（如 UTM 或 World Mollweide）以获得准确的米制单位，或者使用支持球面计算的工具。

### 2.4 压缩
- 为了减小文件体积，WorldPop 的 TIFF 文件通常使用 **LZW** 或 **Deflate** 等无损压缩算法。这意味着在读取时需要解压缩，但不会丢失数据精度。

## 3. 如何读取和处理 WorldPop TIFF

### 使用 GIS 软件
- **QGIS / ArcGIS**: 直接拖入即可显示。软件会自动识别坐标系并将其叠加在地图上。可以通过“识别”工具点击像素查看具体的人口数值。

### 使用编程语言

#### Python (推荐)
使用 `rasterio` 或 `gdal` 库读取：

```python
import rasterio
import numpy as np

# 打开 TIFF 文件
with rasterio.open('ppp_2020.tif') as src:
    # 读取第一波段数据
    data = src.read(1)
    
    # 获取元数据
    transform = src.transform # 仿射变换参数（像素到坐标的映射）
    crs = src.crs             # 坐标系信息
    nodata = src.nodata       # 无效值定义

    # 处理数据：将 NoData 替换为 0 或 NaN
    data = np.where(data == nodata, 0, data)
    
    # 计算总人口
    total_population = np.sum(data)
    print(f"Total Population: {total_population}")
```

#### R 语言
使用 `terra` 或 `raster` 包：

```R
library(terra)

# 读取 TIFF
pop_raster <- rast("ppp_2020.tif")

# 查看基本信息
print(pop_raster)

# 计算总人口
total_pop <- global(pop_raster, "sum", na.rm=TRUE)
print(total_pop)
```

## 4. 常见注意事项

1.  **浮点数精度**: 由于像素值是模型估计的浮点数，直接加总可能会有微小的精度误差，但在人口统计学意义上通常可以忽略。
2.  **网格面积变化**: 在 WGS84 坐标系下，不同纬度的网格单元实际地面面积是不同的（赤道处最大，向两极逐渐减小）。如果是计算人口**密度**或基于面积的分析，务必考虑到这一点，或者使用 WorldPop 提供的“像素面积” (Pixel Area) 辅助数据集进行加权计算。
3.  **对齐问题**: 在进行多时相分析（如比较 2000 年和 2020 年）时，确保两个 TIFF 文件的网格是对齐的（即 Origin 和 Resolution 完全一致），否则需要进行重采样 (Resampling)。
