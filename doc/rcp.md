# Representative Concentration Pathways (RCPs) 详细介绍

## 1. 什么是 RCPs？

RCPs（Representative Concentration Pathways，代表性浓度路径）是一组用于气候模拟的情景输入，核心是给出未来大气成分（温室气体、气溶胶等）的**浓度与辐射强迫**随时间变化的轨迹，从而驱动气候模式（尤其是 CMIP5 阶段的大气海洋耦合模式与地球系统模式）进行未来气候投影。

RCPs 的命名来自其在 2100 年相对于前工业时期的**总辐射强迫（radiative forcing）水平**（单位 \(W/m^2\)）。常见的四条主路径是：
- RCP2.6
- RCP4.5
- RCP6.0
- RCP8.5

## 2. 关键概念：辐射强迫（Radiative Forcing）

辐射强迫可以理解为“地球能量收支被扰动后的净变化量”。当温室气体浓度升高、吸收并再辐射更多长波辐射时，会导致正的辐射强迫；气溶胶等因素也会影响辐射强迫（可能是正也可能是负）。

- 单位：\(W/m^2\)（瓦特每平方米）
- 参照基线：通常以前工业时期（常以 1750 年附近）为基准
- 含义：数值越高，代表对气候系统的“额外加热效应”越强，但温度升幅并不是简单等于辐射强迫值，还取决于气候敏感度、碳循环反馈等

## 3. 四条主 RCP 路径概览

下表给出四条主 RCP 的直观解读（重点是 2100 年辐射强迫水平与大致趋势形态）：

| 路径 | 2100 年辐射强迫目标（约） | 典型趋势形态 | 常见解读 |
| :--- | :--- | :--- | :--- |
| **RCP2.6** | 2.6 \(W/m^2\) | **先峰值后回落**（中世纪附近峰值约 3 \(W/m^2\) 后下降） | 强减排/强缓解路径 |
| **RCP4.5** | 4.5 \(W/m^2\) | **趋于稳定**（不明显超调） | 中等稳定路径 |
| **RCP6.0** | 6.0 \(W/m^2\) | **趋于稳定**（不明显超调） | 较高稳定路径 |
| **RCP8.5** | 8.5 \(W/m^2\) | **持续上升** | 高强迫/高排放路径 |

RCP2.6 在部分文献中也会看到别名 **RCP3-PD**（Peak & Decline），强调其辐射强迫“峰值后下降”的特征。

## 4. RCPs 是怎么被用到气候模型里的？

以 CMIP5 时代为例，RCPs 的典型使用链路可以概括为：

1. 选择一条 RCP（例如 RCP4.5）
2. 该路径提供（或可推导）未来的排放、土地利用变化、短寿命气候污染物与长寿命温室气体浓度等信息
3. 将浓度/排放等驱动输入到气候模式或地球系统模式
4. 模式输出未来的温度、降水、风场、海平面等气候变量
5. 影响评估（IAV：Impacts, Adaptation, Vulnerability）基于这些气候输出进一步估算风险与影响

RCPs 的设计目标之一是让不同研究在同一组“可比的强迫水平”下进行模拟对比，而不是每个研究都用完全不同的一套情景假设。

## 5. RCPs 与 SRES、SSPs 的关系

### 5.1 与 SRES 的关系（更早一代情景）

在 RCPs 之前，IPCC 常用的是 SRES（Special Report on Emissions Scenarios）系列，它主要以“排放情景”为主。RCPs 的一个关键变化是更强调为气候模型提供**更直接的浓度/强迫轨迹**，以适配气候模式实验设计需求。

### 5.2 与 SSPs 的关系（更现代的情景框架）

SSPs（共享社会经济路径）提供的是**社会经济叙事与量化路径**，而 RCPs 更偏向“物理气候系统驱动”（强迫水平）。

在 CMIP6/IPCC AR6 体系中，更常见的命名方式是 **SSP\(x\)-\(y\)**，例如：
- SSP1-2.6
- SSP2-4.5
- SSP5-8.5

其中的 \(y\) 本质上对应“到 2100 年左右的辐射强迫水平”，与 RCP 的数字含义一致；而 \(x\) 则描述社会经济路径（发展模式、能源结构、城市化等）带来的排放与脆弱性差异。

## 6. 常见误解与使用注意

- RCP 的数字不是温升：例如 RCP4.5 的 “4.5” 是 \(W/m^2\) 的辐射强迫水平，不是 “升温 4.5°C”。温升结果需要由气候模型计算。
- RCP 是“路径”而不是单点：同一条 RCP 的核心是 2100 年附近的强迫水平与整个世纪的时间序列形态，而不仅仅是一个终点值。
- RCP 不是完整社会经济叙事：RCPs 自身并不等价于 “未来世界如何发展” 的完整故事；需要与 SSPs 等社会经济框架结合来讲清楚驱动因素与政策含义。

## 7. 参考资料

- van Vuuren, D. P., et al. (2011). The representative concentration pathways: an overview. Climatic Change. https://link.springer.com/article/10.1007/s10584-011-0148-z
- UK Met Office. UKCP18 Guidance: Representative Concentration Pathways (RCPs). https://www.metoffice.gov.uk/binaries/content/assets/metofficegovuk/pdf/research/ukcp/ukcp18-guidance---representative-concentration-pathways.pdf
- Government of Canada (Climate Modelling and Analysis). Representative Concentration Pathways. https://climate-scenarios.canada.ca/?page=scen-rcp
