## soinc
轻量全文检索工具


docker run -p 1491:1491
-v ./config.cfg:/etc/sonic.cfg
-v ./data/store/:/var/lib/sonic/store/
valeriansaliou/sonic:v1.3.2

在配置文件中，确保：
channel.inet设置为0.0.0.0:1491（这可以让 Sonic 从容器外部到达）
store.kv.path设置为/var/lib/sonic/store/kv/（这让 Sonic 可以访问外部 KV 存储目录）
store.fst.path设置为/var/lib/sonic/store/fst/（这让 Sonic 可以访问外部 FST 存储目录）
