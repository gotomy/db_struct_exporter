# boss数据库表结构
## 表名：imports，描述：
### 表结构

|字段名称	|字段类型	|是否可空	|描述	   |
|-----------|-----------|-----------|----------|
|id	|bigint(19)		|NO		|记录主键,自增	   |
|file_name	|varchar(255)		|YES		|文件名	   |
|original_storage_key	|varchar(255)		|YES		|原始文件云存储的key值	   |
|failed_storage_key	|varchar(255)		|YES		|失败,未能导入的云存储的key值	   |
|source	|varchar(1)		|YES		|来源:1 studio, 2 VCG , 3 GETTY	   |
|total	|int(11)		|YES		|上传记录总计	   |
|failed_total	|int(11)		|YES		|错误记录总计	   |
|failed_gross	|int(11)		|YES		|gross	   |
|failed_total_seller_share	|int(11)		|YES		|	   |
|failed_cogs_amount	|int(11)		|YES		|	   |
|successed_total	|int(11)		|YES		|成功上传总计	   |
|successed_gross	|int(11)		|YES		|	   |
|successed_total_seller_share	|int(11)		|YES		|	   |
|successed_cogs_amount	|int(11)		|YES		|	   |
|size	|int(11)		|YES		|文件大小bytes	   |
|processed_by	|bigint(20)		|YES		|操作员编号	   |
|upload_time	|datetime		|YES		|上传时间	   |
|updated_at	|datetime		|YES		|	   |
### 索引结构
### 建表语句
```
CREATE TABLE `imports` (
  `id` bigint(19) NOT NULL AUTO_INCREMENT COMMENT '记录主键,自增',
  `file_name` varchar(255) DEFAULT NULL COMMENT '文件名',
  `original_storage_key` varchar(255) DEFAULT NULL COMMENT '原始文件云存储的key值',
  `failed_storage_key` varchar(255) DEFAULT NULL COMMENT '失败,未能导入的云存储的key值',
  `source` varchar(1) DEFAULT NULL COMMENT '来源:1 studio, 2 VCG , 3 GETTY',
  `total` int(11) DEFAULT NULL COMMENT '上传记录总计',
  `failed_total` int(11) DEFAULT NULL COMMENT '错误记录总计',
  `failed_gross` int(11) DEFAULT NULL COMMENT 'gross',
  `failed_total_seller_share` int(11) DEFAULT NULL,
  `failed_cogs_amount` int(11) DEFAULT NULL,
  `successed_total` int(11) DEFAULT NULL COMMENT '成功上传总计',
  `successed_gross` int(11) DEFAULT NULL,
  `successed_total_seller_share` int(11) DEFAULT NULL,
  `successed_cogs_amount` int(11) DEFAULT NULL,
  `size` int(11) DEFAULT NULL COMMENT '文件大小bytes',
  `processed_by` bigint(20) DEFAULT NULL COMMENT '操作员编号',
  `upload_time` datetime DEFAULT NULL COMMENT '上传时间',
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `file_name` (`file_name`) USING BTREE,
  KEY `upload_time` (`upload_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1000000054 DEFAULT CHARSET=utf8
```
## 表名：imports_error，描述：
### 表结构

|字段名称	|字段类型	|是否可空	|描述	   |
|-----------|-----------|-----------|----------|
|id	|bigint(19)		|NO		|自增主键	   |
|user_id	|bigint(20)		|YES		|	   |
|photo_id	|bigint(20)		|YES		|素材编号	   |
|sheet_index	|int(11)		|YES		|导入源文件sheet序号	   |
|line	|int(11)		|YES		|导入源文件行号	   |
|failed_reason	|varchar(4000)		|YES		|错误原因, 多个原因以 | 分隔	   |
|imports_id	|bigint(19)		|YES		|对应上传记录的主键	   |
|gross	|int(11)		|YES		|未分成前的稿费总额	   |
|seller_share	|int(11)		|YES		| 摄影师分成拿到金额，以美分记。	   |
|photographer_cut	|float		|YES		|摄影师分成百分比	   |
|license_name	|varchar(255)		|YES		|	   |
|exclusive	|varchar(1)		|YES		|0: non_exc; 1:exclusive	   |
|sales_territory	|varchar(255)		|YES		|销售区域	   |
|created_at	|datetime		|YES		|交易时间	   |
|updated_at	|datetime		|YES		|	   |
### 索引结构
### 建表语句
```
CREATE TABLE `imports_error` (
  `id` bigint(19) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `user_id` bigint(20) DEFAULT NULL,
  `photo_id` bigint(20) DEFAULT NULL COMMENT '素材编号',
  `sheet_index` int(11) DEFAULT NULL COMMENT '导入源文件sheet序号',
  `line` int(11) DEFAULT NULL COMMENT '导入源文件行号',
  `failed_reason` varchar(4000) DEFAULT NULL COMMENT '错误原因, 多个原因以 | 分隔',
  `imports_id` bigint(19) DEFAULT NULL COMMENT '对应上传记录的主键',
  `gross` int(11) DEFAULT NULL COMMENT '未分成前的稿费总额',
  `seller_share` int(11) DEFAULT '0' COMMENT ' 摄影师分成拿到金额，以美分记。',
  `photographer_cut` float DEFAULT NULL COMMENT '摄影师分成百分比',
  `license_name` varchar(255) DEFAULT NULL,
  `exclusive` varchar(1) DEFAULT NULL COMMENT '0: non_exc; 1:exclusive',
  `sales_territory` varchar(255) DEFAULT NULL COMMENT '销售区域',
  `created_at` datetime DEFAULT NULL COMMENT '交易时间',
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `imports_id` (`imports_id`) USING BTREE,
  KEY `photo_id` (`photo_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1000017748 DEFAULT CHARSET=utf8
```
## 表名：payments，描述：
### 表结构

|字段名称	|字段类型	|是否可空	|描述	   |
|-----------|-----------|-----------|----------|
|id	|int(11)		|NO		|	   |
|state	|int(11)		|YES		|	   |
|capture_date	|datetime		|YES		|	   |
|created_at	|datetime		|YES		|	   |
|updated_at	|datetime		|YES		|	   |
### 索引结构
### 建表语句
```
CREATE TABLE `payments` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `state` int(11) DEFAULT NULL,
  `capture_date` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `state` (`state`)
) ENGINE=InnoDB AUTO_INCREMENT=1411444 DEFAULT CHARSET=utf8
```
## 表名：payout_requests，描述：
### 表结构

|字段名称	|字段类型	|是否可空	|描述	   |
|-----------|-----------|-----------|----------|
|id	|bigint(19)		|NO		|	   |
|user_id	|bigint(20)		|YES		|	   |
|amount	|int(11)		|YES		|	   |
|method	|varchar(255)		|YES		|	   |
|address_id	|int(11)		|YES		|	   |
|paypal_account	|varchar(255)		|YES		|	   |
|paypal_txn_id	|varchar(255)		|YES		|	   |
|cheque_number	|varchar(255)		|YES		|	   |
|comment	|text		|YES		|	   |
|admin_id	|bigint(20)		|YES		|	   |
|paypal_resource	|text		|YES		|	   |
|has_photo	|tinyint(1)		|YES		|	   |
|paid_date	|datetime		|YES		|	   |
|state	|varchar(255)		|YES		|	   |
|fee	|int(11)		|YES		|	   |
|withholding	|int(11)		|YES		|	   |
|created_at	|datetime		|YES		|	   |
|updated_at	|datetime		|YES		|	   |
### 索引结构
### 建表语句
```
CREATE TABLE `payout_requests` (
  `id` bigint(19) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `amount` int(11) DEFAULT NULL,
  `method` varchar(255) DEFAULT NULL,
  `address_id` int(11) DEFAULT NULL,
  `paypal_account` varchar(255) DEFAULT NULL,
  `paypal_txn_id` varchar(255) DEFAULT NULL,
  `cheque_number` varchar(255) DEFAULT NULL,
  `comment` text,
  `admin_id` bigint(20) DEFAULT NULL,
  `paypal_resource` text,
  `has_photo` tinyint(1) DEFAULT '0',
  `paid_date` datetime DEFAULT NULL,
  `state` varchar(255) DEFAULT NULL,
  `fee` int(11) DEFAULT NULL,
  `withholding` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `index_payout_requests_on_state` (`state`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=46432 DEFAULT CHARSET=utf8
```
## 表名：payout_requests_sales，描述：
### 表结构

|字段名称	|字段类型	|是否可空	|描述	   |
|-----------|-----------|-----------|----------|
|payout_request_id	|bigint(19)		|NO		|	   |
|sales_id	|bigint(19)		|NO		|	   |
### 索引结构
### 建表语句
```
CREATE TABLE `payout_requests_sales` (
  `payout_request_id` bigint(19) NOT NULL DEFAULT '0',
  `sales_id` bigint(19) NOT NULL DEFAULT '0',
  PRIMARY KEY (`payout_request_id`,`sales_id`) USING BTREE,
  KEY `payout_request_id` (`payout_request_id`) USING BTREE,
  KEY `sales_id` (`sales_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8
```
## 表名：sales，描述：
### 表结构

|字段名称	|字段类型	|是否可空	|描述	   |
|-----------|-----------|-----------|----------|
|id	|bigint(19)		|NO		|	   |
|user_id	|bigint(20)		|YES		|	   |
|buyer_id	|bigint(20)		|YES		|Vcg (71576437 & 14466947);Getty (33778897);500px Enterprise (29632435);500px Marketplace (6728138)	   |
|distributor_id	|bigint(20)		|YES		|Real buyer_id, purchase_items表中的user_id	   |
|stock_photo_id	|varchar(255)		|YES		|分销平台素材编号	   |
|asset_type	|varchar(1)		|YES		|1:picture;2:video;3music	   |
|asset_family	|varchar(1)		|YES		|1.editorial;2.creative	   |
|photo_id	|bigint(20)		|YES		|	   |
|payment_id	|int(11)		|YES		|暂时未用到	   |
|state	|varchar(255)		|YES		|1) 不计入余额的状态 1.a) 买家相关 Waiting_for_authorization 等待支付系统授权 Authorization_failed 授权失败 Waiting_for_payment 等待支付系统完成付款 Refunded 已退款 voided 支付系统授权已失效  failed 支付系统处理失败 disputed 信用卡交易被申诉（相当于退款） 1.b) 摄影师打款相关 Submitted_for_payout 已申请支付摄影师稿费 Paid_out 摄影师打款成功  2) 计入当前余额的状态 Ready_for_ship 买家付款完成  Printing 正在打印 shipped 已经发出打印件	   |
|payout_request_id	|int(11)		|YES		|关联的提现记录ID	   |
|gross	|int(11)		|YES		|未分成前的稿费总额	   |
|seller_share	|int(11)		|YES		| 摄影师分成拿到金额，以美分记。	   |
|photographer_cut	|float		|YES		|摄影师分成百分比	   |
|payment_method	|int(11)		|YES		|销售渠道或零售付款方式，见文档详细解释	   |
|license_name	|varchar(255)		|YES		|	   |
|exclusive	|varchar(1)		|YES		|0: non_exc; 1:exclusive	   |
|sales_territory	|varchar(255)		|YES		|销售区域	   |
|industry	|varchar(255)		|YES		|行业	   |
|usage	|longtext		|YES		|用途	   |
|original_sales_id	|int(19)		|YES		|原始销售记录ID	   |
|is_refunded	|varchar(1)		|YES		|是否被退款. 0:否;1是	   |
|vendor_license_id	|varchar(255)		|YES		|销售记录所属的订单编号. 退款时通过invoice_number, user_id, photo_id, gross找到原销售记录.	   |
|cancelled	|tinyint(1)		|YES		|原marketplace取消购买标记, 历史数据导入,保证新系统数据兼容性. 0:否;1:是	   |
|imports_id	|bigint(19)		|YES		|关联imports表的id	   |
|sheet_index	|int(11)		|YES		|对应源文件Sheet序号	   |
|line	|int(11)		|YES		|对应原导入文件中的行号	   |
|source	|varchar(1)		|YES		|来源. 1.studio;2.Getty;3VCG	   |
|created_at	|datetime		|YES		|交易时间	   |
|updated_at	|datetime		|YES		|	   |
### 索引结构
### 建表语句
```
CREATE TABLE `sales` (
  `id` bigint(19) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `buyer_id` bigint(20) DEFAULT NULL COMMENT 'Vcg (71576437 & 14466947);Getty (33778897);500px Enterprise (29632435);500px Marketplace (6728138)',
  `distributor_id` bigint(20) DEFAULT NULL COMMENT 'Real buyer_id, purchase_items表中的user_id',
  `stock_photo_id` varchar(255) DEFAULT NULL COMMENT '分销平台素材编号',
  `asset_type` varchar(1) DEFAULT NULL COMMENT '1:picture;2:video;3music',
  `asset_family` varchar(1) DEFAULT NULL COMMENT '1.editorial;2.creative',
  `photo_id` bigint(20) DEFAULT NULL,
  `payment_id` int(11) DEFAULT NULL COMMENT '暂时未用到',
  `state` varchar(255) DEFAULT NULL COMMENT '1) 不计入余额的状态\n1.a) 买家相关\nWaiting_for_authorization 等待支付系统授权\nAuthorization_failed 授权失败\nWaiting_for_payment 等待支付系统完成付款\nRefunded 已退款\nvoided 支付系统授权已失效 \nfailed 支付系统处理失败\ndisputed 信用卡交易被申诉（相当于退款）\n1.b) 摄影师打款相关\nSubmitted_for_payout 已申请支付摄影师稿费\nPaid_out 摄影师打款成功\n\n2) 计入当前余额的状态\nReady_for_ship 买家付款完成 \nPrinting 正在打印\nshipped 已经发出打印件',
  `payout_request_id` int(11) DEFAULT NULL COMMENT '关联的提现记录ID',
  `gross` int(11) DEFAULT NULL COMMENT '未分成前的稿费总额',
  `seller_share` int(11) DEFAULT '0' COMMENT ' 摄影师分成拿到金额，以美分记。',
  `photographer_cut` float DEFAULT NULL COMMENT '摄影师分成百分比',
  `payment_method` int(11) DEFAULT NULL COMMENT '销售渠道或零售付款方式，见文档详细解释',
  `license_name` varchar(255) DEFAULT NULL,
  `exclusive` varchar(1) DEFAULT NULL COMMENT '0: non_exc; 1:exclusive',
  `sales_territory` varchar(255) DEFAULT NULL COMMENT '销售区域',
  `industry` varchar(255) DEFAULT NULL COMMENT '行业',
  `usage` longtext COMMENT '用途',
  `original_sales_id` int(19) DEFAULT NULL COMMENT '原始销售记录ID',
  `is_refunded` varchar(1) DEFAULT NULL COMMENT '是否被退款. 0:否;1是',
  `vendor_license_id` varchar(255) DEFAULT NULL COMMENT '销售记录所属的订单编号. 退款时通过invoice_number, user_id, photo_id, gross找到原销售记录.',
  `cancelled` tinyint(1) DEFAULT '0' COMMENT '原marketplace取消购买标记, 历史数据导入,保证新系统数据兼容性. 0:否;1:是',
  `imports_id` bigint(19) DEFAULT NULL COMMENT '关联imports表的id',
  `sheet_index` int(11) DEFAULT NULL COMMENT '对应源文件Sheet序号',
  `line` int(11) DEFAULT NULL COMMENT '对应原导入文件中的行号',
  `source` varchar(1) DEFAULT NULL COMMENT '来源. 1.studio;2.Getty;3VCG',
  `created_at` datetime DEFAULT NULL COMMENT '交易时间',
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `state` (`state`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `payout_request_id` (`payout_request_id`) USING BTREE,
  KEY `created_at` (`created_at`) USING BTREE,
  KEY `seller_share` (`seller_share`) USING BTREE,
  KEY `imports_id` (`imports_id`) USING BTREE,
  KEY `buyer_id` (`buyer_id`) USING BTREE,
  KEY `sales_territory` (`sales_territory`) USING BTREE,
  KEY `original_sales_id` (`original_sales_id`),
  KEY `invoice_number` (`vendor_license_id`),
  KEY `photo_id` (`photo_id`),
  KEY `gross` (`gross`),
  KEY `payment_id` (`payment_id`),
  KEY `payment_id_2` (`payment_id`)
) ENGINE=InnoDB AUTO_INCREMENT=561044 DEFAULT CHARSET=utf8
```
## 表名：territory，描述：
### 表结构

|字段名称	|字段类型	|是否可空	|描述	   |
|-----------|-----------|-----------|----------|
|territory_id	|varchar(32)		|NO		|主键id	   |
|territory_name	|varchar(255)		|YES		|适用领地范围名称	   |
|created_stamp	|datetime		|YES		|创建时间	   |
|territory_name_zh	|varchar(255)		|YES		|适用国家名称-中文	   |
|territory_name_en	|varchar(255)		|YES		|适用国家名称-英文	   |
|territory_code	|varchar(20)		|YES		|国家编码	   |
|type	|varchar(60)		|YES		|领地类型	   |
|meta_territory_name_zh	|varchar(255)		|YES		|区域名称-中文，如亚洲	   |
|meta_territory_name_en	|varchar(255)		|YES		|区域名称-英文	   |
|hidden_flag	|varchar(1)		|YES		|作废标志	   |
|country	|int(11)		|YES		|国家数	   |
|created_time	|timestamp		|YES		|	   |
|created_by	|varchar(32)		|YES		|	   |
|updated_time	|timestamp		|YES		|	   |
|updated_by	|varchar(32)		|YES		|	   |
### 索引结构
### 建表语句
```
CREATE TABLE `territory` (
  `territory_id` varchar(32) NOT NULL COMMENT '主键id',
  `territory_name` varchar(255) DEFAULT NULL COMMENT '适用领地范围名称',
  `created_stamp` datetime DEFAULT NULL COMMENT '创建时间',
  `territory_name_zh` varchar(255) DEFAULT NULL COMMENT '适用国家名称-中文',
  `territory_name_en` varchar(255) DEFAULT NULL COMMENT '适用国家名称-英文',
  `territory_code` varchar(20) DEFAULT NULL COMMENT '国家编码',
  `type` varchar(60) DEFAULT NULL COMMENT '领地类型',
  `meta_territory_name_zh` varchar(255) DEFAULT NULL COMMENT '区域名称-中文，如亚洲',
  `meta_territory_name_en` varchar(255) DEFAULT NULL COMMENT '区域名称-英文',
  `hidden_flag` varchar(1) DEFAULT NULL COMMENT '作废标志',
  `country` int(11) DEFAULT NULL COMMENT '国家数',
  `created_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(32) DEFAULT NULL,
  `updated_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`territory_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='区域'
```
## 表名：white_list，描述：
### 表结构

|字段名称	|字段类型	|是否可空	|描述	   |
|-----------|-----------|-----------|----------|
|id	|bigint(19)		|NO		|记录主键，自增	   |
|user_id	|bigint(20)		|NO		|用户id	   |
|authority_name	|varchar(64)		|NO		|权限名	   |
|created_time	|datetime		|NO		|创建时间	   |
### 索引结构
### 建表语句
```
CREATE TABLE `white_list` (
  `id` bigint(19) NOT NULL AUTO_INCREMENT COMMENT '记录主键，自增',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `authority_name` varchar(64) NOT NULL COMMENT '权限名',
  `created_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1000000015 DEFAULT CHARSET=utf8
```