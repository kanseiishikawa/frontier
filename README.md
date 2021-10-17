# frontier βver
特徴量の有効性を自動で検出する

## コマンド
```
cd src/   
go run main.go -target=target.json -fevalue=fevalue.js
```

## コマンドオプション
### -target
デフォルトはtarget.json  
答えデータを保存したjsonのパスを指定する。

### -fevalue
デフォルトはfevalue.json  
特徴量データを保存したjsonのパスを指定する。

## サンプル
sample/の配下にtargetとfevalueのサンプルのjsonが保存されている

### sample_target.json
```
{  
    "Info": [  
        {  
            "Name": "time",  
            "Up": false  
        },  
        {  
            "Name": "rank",  
            "Up": false  
        },  
        {  
            "Name": "up_time",  
            "Up": false  
        }  
    ],  
    "Value": [  
        [  
            107.69999999999999,  
            5.0,  
            30.4  
        ],  
        [  
            107.69999999999999,  
            6.0,  
            30.4  
        ],  
        [  
            107.1,  
            3.0,  
            30.4  
        ]  
    ]  
}
```

### sample_fevalue.json
```
{  
    "Info": [  
        {  
            "Name": "speed"  
        },  
        {  
            "Name": "up_speed"  
        },  
        {  
            "Name": "pace_speed"  
        }  
    ],  
    "Value": [  
        [  
            77.14372039999999,  
            -6.641250619074621,  
            -9.719415395017798  
        ],  
        [  
            58.067769999999996,  
            -11.05551458332204,  
            -15.366911620957517  
        ],  
        [  
            92.0558812,  
            6.33984480748319,  
            3.330188602502667  
        ],  
        [  
            86.95316679999999,  
            4.0445536669587785,  
            1.1708512397027508  
        ]  
    ]  
}
```
