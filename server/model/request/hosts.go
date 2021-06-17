package request

import "quan/model"

type HostsSearch struct{
    model.Hosts
    PageInfo
}