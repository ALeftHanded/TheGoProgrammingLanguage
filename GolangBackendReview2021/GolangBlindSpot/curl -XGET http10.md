```shell
curl -XGET "http://10.131.225.17:9203/supplier_sku_index/_search" -H 'Content-Type: application/json' -d'{"query": {"bool": {"must": [{"term": {"sku_id": "14146574741_142189172714"}}]}}}'


curl -XGET "http://10.131.225.17:9203/supplier_sku_index/_search" -H 'Content-Type: application/json' -d'{"query": {"bool": {"must": [{"term": {"sbs_vendor_id": "SBSSG000495"}}]}}}'
```

```shell
curl -XPOST "https://retail.ssc.shopeemobile.com/api/supplier/sku/call_kafka_handler" -H 'Content-Type: application/json' -d'{
    "type": "SupplierSkuTab",
    "requester": "PMS",
    "request_time": 1643352879,
    "uuid": "12754a75-1675-482c-8ba2-ee284a6e664a",
    "data": [
        {
            "sbs_vendor_id": "SBSPH000680",
            "sku_id": "9881424753_56571823854"
        },
        {
            "sbs_vendor_id": "SBSPH000680",
            "sku_id": "9881424777_101124583096"
        },
        {
            "sbs_vendor_id": "SBSPH000680",
            "sku_id": "5595336492_46571824164"
        },
        {
            "sbs_vendor_id": "SBSPH000680",
            "sku_id": "9881424800_101124583108"
        },
        {
            "sbs_vendor_id": "SBSPH000680",
            "sku_id": "2982014205_111124583341"
        },
        {
            "sbs_vendor_id": "SBSPH000680",
            "sku_id": "9981424800_36571824071"
        },
        {
            "sbs_vendor_id": "SBSPH000680",
            "sku_id": "9881424770_101124582990"
        },
        {
            "sbs_vendor_id": "SBSPH000680",
            "sku_id": "5895336385_26902549998"
        },
        {
            "sbs_vendor_id": "SBSPH000680",
            "sku_id": "5895336381_26902549993"
        },
        {
            "sbs_vendor_id": "SBSPH000680",
            "sku_id": "9981424818_56571824050"
        },
        {
            "sbs_vendor_id": "SBSPH000680",
            "sku_id": "3078819586_101124583527"
        },
        {
            "sbs_vendor_id": "SBSPH000680",
            "sku_id": "9981424662_56571799125"
        }
    ]
}'
```

SBSPH000680

for test



```shell
curl -XPOST "https://retail.ssc.test.shopeemobile.com/api/supplier/sku/call_kafka_handler?csrf_token=QA_IS_AWESOME\!OH_YEAH\!" -H 'Content-Type: application/json' -d'{
    "type": "SupplierSkuTab",
    "requester": "PMS",
    "request_time": 1643352879,
    "uuid": "12754a75-1675-482c-8ba2-ee284a6e664a",
    "data": [
    {
        "sbs_vendor_id": "SBSID000085",
        "sku_id": "3602915156_2019223348"
    },
    {
        "sbs_vendor_id": "SBSID000085",
        "sku_id": "3602915156_2019223348"
    },
    {
        "sbs_vendor_id": "SBSSG000001",
        "sku_id": "1540331_1403229"
    },
    {
        "sbs_vendor_id": "SBSID000085",
        "sku_id": "3300210709_10002001994"
    },
    {
        "sbs_vendor_id": "SBSID000085",
        "sku_id": "3300210709_10002001994"
    },
    {
        "sbs_vendor_id": "SBSID000085",
        "sku_id": "3802907792_10020064573"
    },
    {
        "sbs_vendor_id": "SBSID000085",
        "sku_id": "3802907792_10020064573"
    },
    {
        "sbs_vendor_id": "SBSID000085",
        "sku_id": "3602898227_10019877127"
    },
    {
        "sbs_vendor_id": "SBSID000085",
        "sku_id": "3602898227_10019877127"
    },
    {
        "sbs_vendor_id": "SBSID000085",
        "sku_id": "2702925277_10020172253"
    }
]
}'
```



```shell
curl -XGET "http://10.131.225.17:9203/supplier_sku_index/_search" -H 'Content-Type: application/json' -d '{
    "query": {
        "bool": {
            "must": [
                {
                    "match": {
                        "sbs_vendor_id": "SBSPH000680"
                    }
                },
                {
                    "match": {
                        "sku_id": "5595336492_46571824164"
                    }
                }
            ]
        }
    }
}'
```