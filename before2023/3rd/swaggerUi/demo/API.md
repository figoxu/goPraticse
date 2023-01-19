
# Swagger Example API
Swagger Example API

Table of Contents

1. [Order management API](#orders)
1. [get string by ID](#testapi)

<a name="testapi"></a>

## testapi

| Specification | Value |
|-----|-----|
| Resource Path | /testapi |
| API Version | 1.0.0 |
| BasePath for the API | {{.}} |
| Consumes | application/json |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /testapi/get-string-by-int/\{some_id\} | [GET](#GetStringByInt) | get string by ID |
| /testapi/get-struct-by-int/\{some_id\} | [GET](#GetStructByInt) | get struct by ID |
| /testapi/get-struct2-by-int/\{some_id\} | [GET](#GetStruct2ByInt) | get struct2 by ID |
| /testapi/get-simple-array-by-string/\{some_id\} | [GET](#GetSimpleArrayByString) | get simple array by ID |
| /testapi/get-struct-array-by-string/\{some_id\} | [GET](#GetStructArrayByString) | get struct array by ID |
| /testapi/get-interface | [GET](#GetInterface) | get interface |
| /testapi/get-simple-aliased | [GET](#GetSimpleAliased) | get simple aliases |
| /testapi/get-array-of-interfaces | [GET](#GetArrayOfInterfaces) | get array of interfaces |
| /testapi/get-struct3 | [GET](#GetStruct3) | get struct3 |



<a name="GetStringByInt"></a>

#### API: /testapi/get-string-by-int/\{some_id\} (GET)


get string by ID



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| some_id | path | int | Some ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | string |  |
| 400 | object | [APIError](#mytest.swagger-demo.APIError) | We need ID!! |
| 404 | object | [APIError](#mytest.swagger-demo.APIError) | Can not find ID |


<a name="GetStructByInt"></a>

#### API: /testapi/get-struct-by-int/\{some_id\} (GET)


get struct by ID



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| some_id | path | int | Some ID | Yes |
| offset | query | int | Offset | Yes |
| limit | query | int | Offset | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [StructureWithEmbededStructure](#mytest.swagger-demo.StructureWithEmbededStructure) |  |
| 400 | object | [APIError](#mytest.swagger-demo.APIError) | We need ID!! |
| 404 | object | [APIError](#mytest.swagger-demo.APIError) | Can not find ID |


<a name="GetStruct2ByInt"></a>

#### API: /testapi/get-struct2-by-int/\{some_id\} (GET)


get struct2 by ID



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| some_id | path | int | Some ID | Yes |
| offset | query | int | Offset | Yes |
| limit | query | int | Offset | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [StructureWithEmbededPointer](#mytest.swagger-demo.StructureWithEmbededPointer) |  |
| 400 | object | [APIError](#mytest.swagger-demo.APIError) | We need ID!! |
| 404 | object | [APIError](#mytest.swagger-demo.APIError) | Can not find ID |


<a name="GetSimpleArrayByString"></a>

#### API: /testapi/get-simple-array-by-string/\{some_id\} (GET)


get simple array by ID



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| some_id | path | string | Some ID | Yes |
| offset | query | int | Offset | Yes |
| limit | query | int | Offset | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | array | string |  |
| 400 | object | [APIError](#mytest.swagger-demo.APIError) | We need ID!! |
| 404 | object | [APIError](#mytest.swagger-demo.APIError) | Can not find ID |


<a name="GetStructArrayByString"></a>

#### API: /testapi/get-struct-array-by-string/\{some_id\} (GET)


get struct array by ID



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| some_id | path | string | Some ID | Yes |
| offset | query | int | Offset | Yes |
| limit | query | int | Offset | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | array | [SimpleStructureWithAnnotations](#mytest.swagger-demo.SimpleStructureWithAnnotations) |  |
| 400 | object | [APIError](#mytest.swagger-demo.APIError) | We need ID!! |
| 404 | object | [APIError](#mytest.swagger-demo.APIError) | Can not find ID |


<a name="GetInterface"></a>

#### API: /testapi/get-interface (GET)


get interface



| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [InterfaceType](#mytest.swagger-demo.InterfaceType) |  |
| 400 | object | [APIError](#mytest.swagger-demo.APIError) | We need ID!! |
| 404 | object | [APIError](#mytest.swagger-demo.APIError) | Can not find ID |


<a name="GetSimpleAliased"></a>

#### API: /testapi/get-simple-aliased (GET)


get simple aliases



| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | string |  |
| 400 | object | [APIError](#mytest.swagger-demo.APIError) | We need ID!! |
| 404 | object | [APIError](#mytest.swagger-demo.APIError) | Can not find ID |


<a name="GetArrayOfInterfaces"></a>

#### API: /testapi/get-array-of-interfaces (GET)


get array of interfaces



| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | array | [InterfaceType](#mytest.swagger-demo.InterfaceType) |  |
| 400 | object | [APIError](#mytest.swagger-demo.APIError) | We need ID!! |
| 404 | object | [APIError](#mytest.swagger-demo.APIError) | Can not find ID |


<a name="GetStruct3"></a>

#### API: /testapi/get-struct3 (GET)


get struct3



| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [StructureWithSlice](#mytest.swagger-demo.StructureWithSlice) |  |
| 400 | object | [APIError](#mytest.swagger-demo.APIError) | We need ID!! |
| 404 | object | [APIError](#mytest.swagger-demo.APIError) | Can not find ID |




### Models

<a name="mytest.swagger-demo.APIError"></a>

#### APIError

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| ErrorCode | int |  |
| ErrorMessage | string |  |

<a name="mytest.swagger-demo.InterfaceType"></a>

#### InterfaceType

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|

<a name="mytest.swagger-demo.SimpleStructureWithAnnotations"></a>

#### SimpleStructureWithAnnotations

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| Name | string |  |
| id | int |  |

<a name="mytest.swagger-demo.StructureWithEmbededPointer"></a>

#### StructureWithEmbededPointer

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| Id | int |  |
| Name | array |  |

<a name="mytest.swagger-demo.StructureWithEmbededStructure"></a>

#### StructureWithEmbededStructure

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| Id | int |  |
| Name | array |  |

<a name="mytest.swagger-demo.StructureWithSlice"></a>

#### StructureWithSlice

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| Id | int |  |
| Name | array |  |


