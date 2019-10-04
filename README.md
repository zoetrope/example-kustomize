# Example of Kustomize

Kustomize v3.1.0 の設定例です。

## 使い方

以下で、チェックアウト及び kustomize のバイナリをダウンロードします。

``` console
$ git clone git@github.com:masa213f/example-kustomize.git
$ cd example-kustomize
$ make setup
```

以下でサンプルのビルドができます。

``` console
$ bin/example-build examples/<dir>
```

実行例は以下のとおり。ビルド結果とbaseとの差分が表示されます。

``` console
$ bin/example-build examples/patch-strategicmerge/
=== kustomize version ===
Version: {KustomizeVersion:3.1.0 GitCommit:95f3303493fdea243ae83b767978092396169baf BuildDate:2019-07-26T18:11:16Z GoOs:linux GoArch:amd64}
=== kustomize build ===
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: hoge
  name: deployment1
spec:
  replicas: 999
  selector:
    matchLabels:
      app: hoge
  template:
    metadata:
      labels:
        app: hoge
    spec:
      containers:
      - command:
        - sh
        - -c
        - sleep 3600
        image: example:2019
        name: example-container
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: piyo
  name: deployment2
spec:
  replicas: 3
  selector:
    matchLabels:
      app: piyo
  template:
    metadata:
      labels:
        app: piyo
    spec:
      containers:
      - command:
        - sh
        - -c
        - sleep 3600
        image: debian:latest
        name: example-container
      - command:
        - sh
        - -c
        - sleep 3600
        image: debian-debug:latest
        name: example-container2
---
apiVersion: v1
kind: Pod
metadata:
  labels:
    app: example
  name: pod1
spec:
  containers:
  - command:
    - sh
    - -c
    - echo $(EXAMPLE_VARIABLE1) $(EXAMPLE_VARIABLE2)
    image: ubuntu:latest
    name: example-container
=== diff base -> target ===
--- /dev/fd/63  2019-10-04 02:23:06.296167257 +0000
+++ /dev/fd/62  2019-10-04 02:23:06.296167257 +0000
@@ -5,7 +5,7 @@
     app: hoge
   name: deployment1
 spec:
-  replicas: 3
+  replicas: 999
   selector:
     matchLabels:
       app: hoge
@@ -19,14 +19,8 @@
         - sh
         - -c
         - sleep 3600
-        image: ubuntu:latest
+        image: example:2019
         name: example-container
-      - command:
-        - sh
-        - -c
-        - sleep 3600
-        image: ubuntu-debug:latest
-        name: example-container2
 ---
 apiVersion: apps/v1
 kind: Deployment
```
