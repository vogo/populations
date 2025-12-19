# WorldPop 结构详细介绍

## 1. WorldPop 是什么？

WorldPop 是一个开放的空间人口与人口统计数据项目/研究团队与数据平台，目标是生产、维护并开放发布高分辨率的网格化人口与人口学相关数据，用于公共卫生、发展、灾害风险、城市规划、气候影响评估等场景。其核心产出是：把行政区尺度的人口统计（以及其他人口学指标）通过地理空间协变量与统计/机器学习方法，重建为规则网格上的人口估计值（例如 100m 或 1km）。WorldPop 自身将其定位为一个基于大学的跨学科应用研究团队/研究计划，并以开放数据形式发布成果：https://www.worldpop.org/ 以及 https://www.worldpop.org/about/

## 1.1 WorldPop 是一家什么样的机构？

从机构属性来看，WorldPop 更接近“高校研究计划/研究团队”，而不是商业公司。

- WorldPop 在官网介绍中将自身描述为 interdisciplinary applied research group（跨学科应用研究团队），使命是让决策者更好地使用空间人口数据，并推动“everywhere is counted in decision making”。https://www.worldpop.org/
- WorldPop 在 About 页面中明确其 research programme based at the University of Southampton（基于南安普顿大学的研究计划），并强调与政府、联合国机构、捐助方等合作开展共研与能力建设。https://www.worldpop.org/about/
- 南安普顿大学的研究组页面也将 WorldPop 作为大学研究团队进行展示。https://www.southampton.ac.uk/research/groups/worldpop

## 1.2 WorldPop 是公益的还是商业机构？

从公开信息与常见运作方式来看，WorldPop 更符合“学术研究与公共数据基础设施”的属性：

- 其数据以开放访问（open access）为主要特征，并强调透明方法与可复用工具（open data / transparent approaches）。https://www.worldpop.org/ 以及 https://www.worldpop.org/about/
- 其运作通常由科研经费与捐助方资助项目驱动，并通过与政府、联合国机构、大学等合作落地应用。https://www.worldpop.org/about/

在工程语境里，可以把它理解为“由高校研究团队主导、面向公共利益发布开放数据与方法的项目体系”，而不是以售卖数据为核心的商业机构。

## 1.3 WorldPop 如何驱动的？（资金、合作与产出机制）

WorldPop 的驱动机制大体可以概括为“项目制资金 + 多方合作 + 开放发布”：

- **项目制资金**：以科研项目与捐助项目为主，围绕特定主题（例如人口分布、年龄性别结构、城市变化、健康相关指标、未来情景等）持续迭代数据与方法。https://www.worldpop.org/about/
- **多方共研与共建**：与各国政府、UN 机构、捐助方、大学合作开展共研项目，包含数据生产、方法改进与能力建设。https://www.worldpop.org/ 以及 https://www.worldpop.org/about/
- **开放产出与反馈迭代**：强调开放数据与用户反馈，通过持续发布与用户反馈改进数据质量与可用性。https://www.worldpop.org/about/

## 1.4 WorldPop 在什么国家？覆盖哪些地区？

- **机构所在地/依托单位**：WorldPop 表述为基于英国的南安普顿大学（University of Southampton）。https://www.worldpop.org/about/
- **工作范围**：其核心工作面向全球，重点关注需要更细粒度人口数据支撑决策的地区与议题，并通过与政府和国际机构合作在不同国家落地。https://www.worldpop.org/

## 2. WorldPop 的“结构”可以怎么理解？

从使用者角度，WorldPop 的结构通常可以分成三层：

1. **数据产品层（What）**：发布哪些数据类别（人口、年龄性别、出生、城市变化、协变量等）。
2. **建模体系层（How）**：不同数据类别背后的建模路线（Top-down / Bottom-up、Constrained / Unconstrained、是否 UN 调整等）。
3. **分发与访问层（Where/How to access）**：通过 WorldPop Hub 下载、通过 WOPR 获取定制产品、通过 API/应用/第三方平台访问等。

下面按这三层展开。

## 3. 数据产品层：WorldPop 常见数据类别

WorldPop 的数据并不只是一种“人口栅格”。常见结构是“以人口为核心 + 若干派生人口学/发展指标 + 支撑协变量与工具”。在 WorldPop Hub 中可以看到多类数据入口（例如 Population Counts、Age and sex structures、Births、Pregnancies、Urban change 等）。其中与本项目（人口栅格与 API 查询）关联度最高的是前两类。参考：WorldPop Hub 数据类型页。https://hub.worldpop.org/project/list

### 3.1 Population Counts（人口数）

这是最基础的产品：每个栅格单元（像素）的值表示该格内估计的居住人口数量（People per pixel / People per grid cell）。

典型衍生：
- **Population Density（人口密度）**：从人口数栅格除以像素地表面积得到（单位常见为 people/\(km^2\)）。
- **不同分辨率版本**：例如 100m、1km；或以弧秒（3 arc-seconds、30 arc-seconds）表示的近似分辨率。

### 3.2 Age and Sex Structures（年龄与性别结构）

在“总人口”之外，WorldPop 也提供按年龄段与性别拆分的网格人口估计，例如每个像素里 0–1 岁男性、1–4 岁女性等。第三方平台（如 Google Earth Engine）也能看到该类数据的描述与变量列表，并明确其生成思路是基于人口普查与行政区数据的机器学习/达西米特里克重分配（Random Forest-based dasymetric redistribution）。参考：Google Earth Engine 的 WorldPop 年龄性别数据条目。https://developers.google.com/earth-engine/datasets/catalog/WorldPop_GP_100m_pop_age_sex

### 3.3 其他人口学与主题类产品（概览）

WorldPop Hub 还发布多类与人口学/健康/发展相关的派生数据与专题数据，例如：
- Births（出生）
- Pregnancies（怀孕）
- Urban change（城市变化）
- 以及更多专题与协变量数据（用于支撑建模与下游分析）

在实际工程落地里，通常会优先选择“人口数/密度/年龄性别”作为基础图层，再按业务需要叠加其他专题。

## 4. 建模体系层：Top-down / Bottom-up 与 Constrained / Unconstrained

WorldPop 的不同产品在“建模路线”上有重要差异，这也是“结构”里最容易被忽略、但对数据解读最关键的一层。

### 4.1 Top-down vs Bottom-up（两条建模路线）

WorldPop 在方法说明中把数据分为 Top-down 与 Bottom-up 两大类，并给出各自适用场景与优缺点：Top-down 依赖行政区的人口普查/官方统计，并利用高分辨率协变量把人口拆到网格；Bottom-up 更强调在普查缺失/不可靠时利用抽样与统计模型推断，并显式表达不确定性。参考：WorldPop 方法说明页。https://www.worldpop.org/methods/populations/

工程使用建议：
- **做全球一致、多年份对比**：通常倾向 Top-down 的“全球一致方法”产品。
- **做单一国家、单一时间点、追求本地精度**：优先看是否有 WOPR 的“定制/合作生产”结果（见下一节）。

### 4.2 Constrained vs Unconstrained（约束与非约束）

在 Top-down 产品里，常见的进一步区分是：
- **Constrained（约束）**：栅格人口在汇总到某个统计层级时，会被约束去匹配输入的人口总量（例如匹配行政区官方统计、或匹配联合国国家总量）。对“可解释的汇总一致性”更友好。
- **Unconstrained（非约束）**：模型更自由地根据协变量分布估计人口空间分布，汇总后不一定严格等于某个官方总量。对“纯空间分布建模”更直接，但在与官方口径对齐时要更谨慎。

在 WorldPop 的数据介绍中，经常能看到“constrained vs unconstrained”的选择提示。参考：Google Earth Engine 的 constrained 年龄性别数据条目。https://developers.google.com/earth-engine/datasets/catalog/WorldPop_GP_100m_pop_age_sex_cons_unadj

### 4.3 UN-adjusted（联合国总量调整）

对部分产品，WorldPop 还提供将国家总人口调整为匹配联合国人口估计的版本（常见描述为 “UN adjusted”），使跨国对比在国家总量口径上更一致。

## 5. 分发与访问层：WorldPop Hub、WOPR、API 与第三方平台

### 5.1 WorldPop Hub（核心下载入口）

WorldPop Hub（hub.worldpop.org）是面向数据下载与分类浏览的主要入口，按数据类别与版本组织下载。对于需要系统性“批量下载 + 本地处理 + 构建服务”的项目，通常以 Hub 的 GeoTIFF 产品为主。

参考：WorldPop Hub 数据类型页。https://hub.worldpop.org/project/list

### 5.2 WOPR（WorldPop Open Population Repository，定制/试验结果仓库）

WOPR 更像是“面向单国的定制建模成果与相关产物的开放仓库”，强调 bespoke（定制）方法、单一时间点、高分辨率，并允许发布试验性结果。WOPR 主页也明确指出：一致的全球多年数据集仍然从 WorldPop 网站/Hub 获取。参考：WOPR 主页。https://wopr.worldpop.org/

如果你的目标国家有 WOPR 产品，一般优先考虑 WOPR 的定制结果（更可能利用本地数据与更适配的协变量）。

### 5.3 API 与应用

WorldPop 提供面向数据访问的 API 与可视化应用（例如 woprVision、WOPR REST API 等），用于交互式查询与系统对接。参考：WorldPop 官网首页与 API 入口信息。https://www.worldpop.org/ 以及 https://hub.worldpop.org/

### 5.4 第三方平台（例如 Google Earth Engine、HDX）

部分 WorldPop 数据在第三方平台提供镜像或索引，优势是集成生态完善、便于在线计算与可视化，但版本更新与元数据口径需要核对。Google Earth Engine 的数据条目通常会提供分辨率、波段说明、生成方法摘要等信息（见前文链接）。

## 6. 与本仓库的关系：为什么理解 WorldPop 结构很重要？

本仓库的目标是“打包未来人口 TIFF 并提供 REST API 查询人口”。对于这种使用方式，理解 WorldPop 的结构会直接影响工程实现与结果解释：

- **选择哪类产品作为底图**：人口数（PPP） vs 人口密度（PD），以及是否需要年龄/性别拆分。
- **选择约束口径**：是否需要与官方/联合国总量严格对齐，影响你在统计汇总与报告时的口径一致性。
- **分辨率与性能权衡**：100m 数据精度更高但体量更大，1km 更适合全球或大范围在线服务。
- **跨版本一致性**：Global1/Global2/WOPR 的版本与方法差异可能导致跨年/跨国比较出现系统偏差，需要在文档或接口元数据中明确数据来源与版本。

## 7. 参考链接

- WorldPop 官网：https://www.worldpop.org/
- WorldPop Hub（数据类型与下载）：https://hub.worldpop.org/project/list
- WorldPop 方法说明（Top-down/Bottom-up）：https://www.worldpop.org/methods/populations/
- WOPR（WorldPop Open Population Repository）：https://wopr.worldpop.org/
- Google Earth Engine：WorldPop 年龄性别数据说明：https://developers.google.com/earth-engine/datasets/catalog/WorldPop_GP_100m_pop_age_sex
- Google Earth Engine：WorldPop constrained 年龄性别数据说明：https://developers.google.com/earth-engine/datasets/catalog/WorldPop_GP_100m_pop_age_sex_cons_unadj
