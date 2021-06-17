package request

import "quan/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}