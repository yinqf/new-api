package common

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/iancoleman/orderedmap"
)

func ReplaceModelForStream(c *gin.Context, sourceModel string, model string, data string) string {
	// 模型名称不一致，存在模型映射转换
	if sourceModel != model {
		oMap := orderedmap.New()
		err := oMap.UnmarshalJSON([]byte(data))
		if err != nil {
			LogError(c, "模型还原，json解析失败："+err.Error())
		} else {
			oMap.Set("model", sourceModel)
			modifiedJSON, err := json.Marshal(oMap)
			if err != nil {
				LogError(c, "模型还原，json组装失败："+err.Error())
			} else {
				data = string(modifiedJSON)
			}
		}
	}
	return data
}

func ReplaceModelForNoStream(c *gin.Context, sourceModel string, model string, responseBody []byte) []byte {
	// 模型名称不一致，存在模型映射转换
	if sourceModel != model {
		oMap := orderedmap.New()
		err := oMap.UnmarshalJSON(responseBody)
		if err != nil {
			LogError(c, "模型还原，json解析失败："+err.Error())
		} else {
			oMap.Set("model", sourceModel)
			modifiedJSON, err := json.MarshalIndent(oMap, "", "    ")
			if err != nil {
				LogError(c, "模型还原，json组装失败："+err.Error())
			} else {
				responseBody = modifiedJSON
			}
		}
	}
	return responseBody
}
