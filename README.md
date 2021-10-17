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
