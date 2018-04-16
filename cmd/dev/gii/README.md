GII
==============


由于go中目前没有好用的类似yii-gii这样的代码生成工具所以写了一个小服务

这样gii可以通过传递表名来查询对应的go类型  

yii-gii 目前也只有 数据库字段类型 ==》 php-type 的转换 并没有go类型的转换规则 
要自己写这些类型转换是比较费劲的  所以考虑使用已有的优良库来提供功能( 感谢xorm的作者 )

## 使用：
~~~http

    GET  http://localhost:1323/gii/table/tbl_brand
	
	返回：
	
	{
    "content": {
        "Name": "content",
        "TableName": "",
        "FieldName": "",
        "SQLType": {
            "Name": "TEXT",
            "DefaultLength": 0,
            "DefaultLength2": 0
        },
        "IsJSON": false,
        "Length": 0,
        "Length2": 0,
        "Nullable": true,
        "Default": "",
        "Indexes": {},
        "IsPrimaryKey": false,
        "IsAutoIncrement": false,
        "MapType": 0,
        "IsCreated": false,
        "IsUpdated": false,
        "IsDeleted": false,
        "IsCascade": false,
        "IsVersion": false,
        "DefaultIsEmpty": false,
        "EnumOptions": null,
        "SetOptions": null,
        "DisableTimeZone": false,
        "TimeZone": null,
        "Comment": "品牌内容 简介",
        "GoType": "string"
    },
    "desc_apply2goods": {
        "Name": "desc_apply2goods",
        "TableName": "",
        "FieldName": "",
        "SQLType": {
            "Name": "TINYINT",
            "DefaultLength": 1,
            "DefaultLength2": 0
        },
        "IsJSON": false,
        "Length": 1,
        "Length2": 0,
        "Nullable": false,
        "Default": "0",
        "Indexes": {},
        "IsPrimaryKey": false,
        "IsAutoIncrement": false,
        "MapType": 0,
        "IsCreated": false,
        "IsUpdated": false,
        "IsDeleted": false,
        "IsCascade": false,
        "IsVersion": false,
        "DefaultIsEmpty": false,
        "EnumOptions": null,
        "SetOptions": null,
        "DisableTimeZone": false,
        "TimeZone": null,
        "Comment": "简介应用到商品",
        "GoType": "int"
    },
    "display_order": {
        "Name": "display_order",
        "TableName": "",
        "FieldName": "",
        "SQLType": {
            "Name": "INT",
            "DefaultLength": 5,
            "DefaultLength2": 0
        },
        "IsJSON": false,
        "Length": 5,
        "Length2": 0,
        "Nullable": true,
        "Default": "1",
        "Indexes": {},
        "IsPrimaryKey": false,
        "IsAutoIncrement": false,
        "MapType": 0,
        "IsCreated": false,
        "IsUpdated": false,
        "IsDeleted": false,
        "IsCascade": false,
        "IsVersion": false,
        "DefaultIsEmpty": false,
        "EnumOptions": null,
        "SetOptions": null,
        "DisableTimeZone": false,
        "TimeZone": null,
        "Comment": "显示顺序",
        "GoType": "int"
    },
	...
	}
~~~

这样yii端的gii 有了这个映射 提取**GoType** 就可以生成代码啦！

## 注意事项

开发的时候  *config/app.toml* 里面配置还需要针对哪个数据库