# 索引设计

* Minecraft 的音效并不算多，至少目前是**不到万级**的。

## 0. 每个音效要存储的信息

1. **标签列表**（和标签倒排索引相呼应）。
2. 音效名字。
3. 音效真实路径。
4. 下载次数。
5. **音效翻译名列表**。（可能有多个音效翻译名）

## 1. 标签

标签搜索应当采用**倒排索引**。  

```javascript
// Inverted Index
{
    tag1: [sound_hash1, sound_hash2, ...],
    tag2: [sound_hash3, sound_hash4, ...],
    ...
}
```  

## 2. 音效哈希值检索

直接用音效哈希值进行检索，如 `b62ca8ec10d07e6bf5ac8dae0c8c1d2e6a1e3356`。  

## 3. 路径词

比如音效文件是 `minecraft/sounds/ambient/cave/cave18.ogg`，可以通过 `ambient`, `cave`, `cave18`, `cave18.ogg` ，甚至直接输入 `cave/cave18.ogg` 这种路径片段进行搜索。  

一样可以采用**倒排索引**：  

```javascript
{
    "ambient": [sound_id1, sound_id2, ...],
    "cave": [sound_id3, sound_id4, ...],
    "cave18": [sound_id5],
    ...
}
```

如果倒排索引中没有找到，就去前缀树找：  

```javascript
{
    "a": {
        "m": {
            "b": {
                "i": {
                    "e": {
                        "n": {
                            // 标记是否是一个倒排索引项
                            // 如果是，程序就会重新回到倒排索引中查找
                            "t": true,
                        }
                    }
                }
            }
        }
    },
    "c": {
        "a": {
            "v": {
                "e": [sound_id3, sound_id4, ...]
            }
        }
    },
}
```



## 4. 联合搜索

用户可以对 1, 3 进行联合搜索，联合搜索的时候**取交集**即可。


## 5. 模糊搜索


