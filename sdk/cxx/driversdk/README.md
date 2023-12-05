# driversdk
driversdk 实现程序

TODO: 还没测试，代码还需要完善

## 构建
```shell
ARCH=`uname -m`
rm -rf ./build-$ARCH
rm -rf ./install-$ARCH
mkdir build-$ARCH
cd build-$ARCH
cmake .. -DCMAKE_INSTALL_PREFIX=../install-$ARCH
make -j4
make install
```

* 输出目录为 PROJECT_DIR/output-${CMAKE\_SYSTEM\_PROCESSOR}
* 最终输出文件为 

```
├── include
│     └── src
│         ├── context
│         │     └── context.h
│         ├── models
│         │     ├── accessconfig.h
│         │     ├── accesstemplate.h
│         │     ├── arraytype.h
│         │     ├── deviceinfo.h
│         │     ├── deviceproperty.h
│         │     ├── enumtype.h
│         │     ├── enumvalue.h
│         │     ├── event.h
│         │     ├── modelmapping.h
│         │     ├── objecttype.h
│         │     ├── propertyvisitor.h
│         │     └── subdeviceyaml.h
│         ├── plugin
│         │     ├── grpcdriver.h
│         │     ├── grpcreport.h
│         │     ├── idriver.h
│         │     ├── ireport.h
│         │     └── serve.h
│         └── utils
│             └── path.h
└── lib
    ├── libdrivercontext.a
    ├── libdrivermodel.a
    ├── libdriverplugin.a
    └── libdriverutil.a

```

