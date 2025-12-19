# FuturePop 详细介绍

## 1. 什么是 FuturePop？

FuturePop 是由 **WorldPop** 团队开发的一套全球高分辨率未来人口预测数据集。它旨在填补气候变化影响研究中对高精度未来人口数据需求的空白。

传统的全球人口预测（如联合国的《世界人口展望》）通常是国家或区域层面的汇总数据，或者是分辨率较粗（如 1km 或更粗）的网格数据。FuturePop 提供了 **~90米 (3弧秒)** 的超高分辨率全球网格化人口预测，能够更精细地反映未来人口的空间分布变化。

## 2. 核心特征

- **高空间分辨率**：提供约 90m (在赤道处) 的网格数据，远高于许多现有的全球预测数据集。
- **基于 SSP 情景**：数据完全对齐 **共享社会经济路径 (SSPs)** 的五种情景 (SSP1 - SSP5)，确保了与 IPCC 气候变化研究框架的一致性。
- **时间跨度**：覆盖 2020 年至 2100 年，通常以 10 年（或更细）为间隔。
- **全球覆盖**：覆盖全球主要陆地表面（除了南极洲等无人居住区）。

## 3. 方法论

FuturePop 的生成结合了人口学模型和地理空间数据科学方法：

1.  **宏观预测**：使用 SSP 数据库中的国家级人口总量和城市化率预测作为约束条件。
2.  **空间重分布**：
    -   利用机器学习算法（如随机森林）分析当前人口分布与地理协变量（如土地覆盖、夜间灯光、地形、路网距离、旅行时间等）的关系。
    -   引入 **城市扩张模型**，预测未来城市区域的增长。
    -   根据 SSP 情景叙事（例如，SSP1 倾向于紧凑型城市发展，SSP3 倾向于无序扩张），调整城乡人口分布权重。
3.  **自下而上与自上而下结合**：既考虑了局部的空间吸引力因素，又严格遵守国家总量约束。

## 4. 应用场景

FuturePop 数据的出现对于精细尺度的风险评估至关重要，特别是在灾害风险领域：

-   **洪水风险评估**：洪水灾害通常发生在局部尺度（几十米到几百米）。使用粗分辨率数据可能低估或高估受灾人口。FuturePop 能够更准确地识别位于洪泛区内的未来人口。
-   **公共卫生规划**：预测未来疾病传播风险、医疗资源需求分布。
-   **城市规划与基础设施**：为未来的交通、住房和基础设施建设提供数据支持，特别是在快速城市化的发展中国家。
-   **气候适应策略**：帮助决策者识别未来气候脆弱性热点区域（如沿海低洼地区、干旱区）。

## 5. 数据来源与开发团队

-   **开发团队**：由英国布里斯托大学 (University of Bristol) 和南安普顿大学 (University of Southampton) 的 WorldPop 团队合作开发。
-   **资助方**：Wellcome Trust 等机构。
-   **关联项目**：该项目是 WorldPop "Global 2" 计划的一部分，旨在更新和提升全球网格化人口数据的质量。

## 6. 与 WorldPop 其他数据的区别

-   **WorldPop Global 1 (2000-2020)**: 基于历史和当前的人口估计，主要用于回顾性分析和现状评估。
-   **FuturePop**: 专注于 **未来 (2020-2100)** 的预测，基于不同的假设情景 (SSPs)，包含不确定性，适用于前瞻性研究。

## 7. 数据下载方式

FuturePop 的公开下载入口通常有两类：项目主页的下载入口，以及 WorldPop 的公开数据仓库（直链下载）。

### 7.1 项目主页（入口聚合）

- 项目主页：https://www.worldpop.org/futurepop/
- 该页面会汇总当前公开发布的数据产品、版本说明与下载入口（可能随版本更新而变化）。

### 7.2 WorldPop 公开数据仓库（直链下载）

WorldPop 会将部分 FuturePop 相关产品发布在公开数据仓库 `data.worldpop.org` 下。以 “Global 1km-grid population projections … version 0.2” 为例，其版本说明（Release Statement）在：

- https://data.worldpop.org/repo/prj/FuturePop/SSPs_1km_v0_2/Release_Statement_FP_SSPs_1km_v0_2.pdf

该版本通常以按 SSP 场景打包的 ZIP 形式提供（例如 `FuturePop_SSP1_1km_v0_2.zip`），解压后包含多个年份的 GeoTIFF。

下载方式示例（以 SSP1 为例）：

```bash
curl -L -o FuturePop_SSP1_1km_v0_2.zip \
  https://data.worldpop.org/repo/prj/FuturePop/SSPs_1km_v0_2/FuturePop_SSP1_1km_v0_2.zip

unzip FuturePop_SSP1_1km_v0_2.zip -d FuturePop_SSP1_1km_v0_2
```

如果你需要其它情景，将 URL 中的 `SSP1` 替换为 `SSP2`/`SSP3`/`SSP4`/`SSP5` 即可。

## 8. 数据使用方式（通用）

FuturePop/WorldPop 的 GeoTIFF 一般为单波段栅格，像素值表示该网格单元内的估计人口数（常见为浮点值）。常见使用方式：

### 8.1 GIS 软件

- QGIS/ArcGIS 直接加载 `.tif` 即可浏览与取值。
- 常见检查项：坐标系是否为 `EPSG:4326`、分辨率与网格是否符合预期、NoData 区域是否正确显示。

### 8.2 命令行快速检查（GDAL）

```bash
gdalinfo your_futurepop.tif | head -n 60
```

建议至少确认：

- `Coordinate System is: GEOGCS["WGS 84"...]`
- 存在 `GeoTransform`（用于经纬度到像素的映射）
- `Band 1` 为目标人口波段

## 9. 在本仓库中使用 FuturePop（API 查询）

本仓库会加载一个 GeoTIFF 文件，并提供按经纬度 + 半径（米）估算圈内人口的 HTTP API。

### 9.1 选择一个要加载的 GeoTIFF

- 选择一个具体 “年份 + SSP 情景” 的 `.tif` 文件作为底图（例如从上面的 ZIP 解压后选取其中一个年份文件）。
- 该文件需要包含有效地理变换参数（GeoTransform）；否则无法进行经纬度到像素坐标的定位。

### 9.2 启动服务

```bash
export FUTURE_POP_TIFF_PATH=/absolute/path/to/your_futurepop.tif
go run cmd/server/main.go
```

可选端口：

```bash
export SERVER_PORT=:8080
```

### 9.3 查询接口

接口为 `GET /api/v1/population`，参数：

- `lat`：纬度（-90 到 90）
- `lng`：经度（-180 到 180）
- `radius`：半径（米，可选，默认 3000）

示例：

```bash
curl "http://localhost:8080/api/v1/population?lat=39.9042&lng=116.4074&radius=3000"
```

### 9.4 重要说明

- 当前实现将像素值 `< 0` 视为无效（NoData）并忽略；如果你的数据集使用其它 NoData 编码方式，结果可能需要相应调整。
- 半径换算为像素时使用了基于纬度的简化近似（将 1 度纬度近似为 111.32km），适合一般近似查询；对高精度面积/距离分析建议在投影坐标系下进行更严谨计算。
