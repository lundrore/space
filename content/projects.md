# 项目

### 系列项目

[tiny](https://github.com/wasuppu/tiny): [ssloy](https://github.com/ssloy) 的一系列小项目 (tinyrenderer, tinyraytracer, tinyraycaster, tinykaboom, tinycompiler) 的 Go 实现，作者没有使用外部库，因此跟着实现了 obj, tga 文件的读取，数学库用于之后大多数的图形学项目上。 还包括了 [implementing-a-tiny-cpu-rasterizer](https://lisyarus.github.io/blog/posts/implementing-a-tiny-cpu-rasterizer.html) 这个 tinyrasterizer，作者文章只写到了 Part 6，尽管代码仓库进展多一些，还是只跟到了这里

[eva](https://github.com/wasuppu/eva): [Dmitry Soshnikov](https://github.com/DmitrySoshnikov) 的系列课程 (Building an Interpreter from Scratch, Building a Parser from Scratch, Building a Transpiler from Scratch, Building a Typechecker from Scratch) 的 Python 实现，Javascript -> Python, JS 的代码写的太灵活，用 Go 实现的有些痛苦，就全部使用 Python 来做了

[pikuma](https://github.com/wasuppu/pikuma): [gustavopezzi](https://github.com/gustavopezzi) 的系列课程 (3D Computer Graphics Programming, Game Physics Engine Programming, 2D Game Engine Development) 的 Go 实现，游戏引擎部分作者嵌入了 lua 解释器，这一步之后没有向后做

[inoneweekend](https://github.com/wasuppu/inoneweekend): [Ray Tracing in One Weekend](https://raytracing.github.io/) 系列 Go 实现，还有文章 [Rasterization in One Weekend](https://tayfunkayhan.wordpress.com/2018/11/24/rasterization-in-one-weekend-part-i/) 的前两部分，用 tiny 中的数学库来代替 glm，glm 的行为有些不同，做了一些修改

### 书籍代码

[computation](https://github.com/wasuppu/computation): [computationbook](https://computationbook.com/) 中译《计算的本质——深入剖析程序和计算机》，关于计算理论介绍的书籍，Ruby -> Go

[dsacc](https://github.com/wasuppu/dsacc): [数据结构与算法分析——C语言描述](https://book.douban.com/subject/1139426/) 数据结构和算法，C -> Go

[groundup](https://github.com/wasuppu/groundup): [Ray Tracing from the Ground Up](https://book.douban.com/subject/2584519/) 实现小型的光线追踪器，C++ -> Go

[baseline](https://github.com/wasuppu/baseline): [compiling to assembly from scratch](https://keleshev.com/compiling-to-assembly-from-scratch/) 编译器，书中解析是由解析组合子完成，完成第一部分，Typescript -> Go

[dezero](https://github.com/wasuppu/dezero): [深度学习入门2——自制框架](https://book.douban.com/subject/36303408/) Python -> Go，完成到步骤43 神经网络，numpy 的 broadcasting 使复刻的有些别扭

### 从头开始加参考

基本是从头开始写，代码结构不同或是没有参考代码

[bmp](https://github.com/wasuppu/bmp): 一个 bmp 图片格式的解析器，结构体定义和注释主要来自 [[MS-WMF]: Windows Metafile Format](http://msdn.microsoft.com/en-us/library/cc250370.aspx)，

[elf](https://github.com/wasuppu/elf): linux 下可执行文件 elf 的解析器，主要参照文档 [TIS1.1.pdf](https://refspecs.linuxfoundation.org/elf/TIS1.1.pdf) 实现，
借鉴了 [finixbit/elf-parser](https://github.com/finixbit/elf-parser/) 和 [macmade/ELFDump](https://github.com/macmade/ELFDump/)，输出样式仿照了 readelf，实现了其 -h, -l, -S, -s 这几个选项

[parc](https://github.com/wasuppu/parc): OCaml 写的解析组合子，主要参照 [Graham Hutton](https://people.cs.nott.ac.uk/pszgmh/) 书《Programming in Haskell》的第十三章以及作者相关的两篇论文，
另外用该解析组合子实现 [wyas](https://en.wikibooks.org/wiki/Write_Yourself_a_Scheme_in_48_Hours)，Towards a Standard Library 一章节实现的有些问题

[sexp](https://github.com/wasuppu/sexp): 一个简单的 s 表达式的解析，打算用于 [ucsd compilers s23](https://ucsd-compilers-s23.github.io/index.html)，参照了项目使用的 Rust 的 [sexp](https://crates.io/crates/sexp) 以及 eva 中用到的 s表达式的解析，框架还是来自于 [interpreterbook](https://interpreterbook.com/)

[codecrafters](https://github.com/wasuppu/codecrafters): [codecrafters](https://app.codecrafters.io/catalog) 网站的挑战解，代码没有整理，完成的 shell 和 browser 网站又有更新，有时间继续完成

[bencode](https://github.com/wasuppu/bencode): 为写 codercrafers 的 bitorrent 做的准备，将前面的手动解析的代码用 marshal/unmarshal 进行包装，方便使用

[argparse](https://github.com/wasuppu/argparse): 刚开始学 Go 写的简单的命令行参数解析器，在 [argparse-rosetta-rs](https://github.com/rosetta-rs/argparse-rosetta-rs) 发现的 [pico-args](https://github.com/razrfalcon/pico-args) 觉得代码量不大，就自己随便开始写了，逻辑混乱，乱七八糟，参考了 go 的 flag 和 python 的 getopt

### 重写加上修改

有参照物的代码，但是进行了修改，或是结合多个参考的混合物

[automata](https://github.com/wasuppu/automata): 自动机是来自 [Dmitry Soshnikov](https://github.com/DmitrySoshnikov) 的自动机课程，解析部分参照 [implementing a regular expression engine](https://deniskyashif.com/2019/02/17/implementing-a-regular-expression-engine/) 结合自己代码生成的 AST

[lept](https://github.com/wasuppu/lept): [miloyip](https://github.com/miloyip) 的 [json-tutorial](https://github.com/miloyip/json-tutorial)，依照其步骤，但是采取 [interpreterbook](https://interpreterbook.com/) 的代码框架进行解析的

[chibiccgo](https://github.com/wasuppu/chibiccgo): [rui314](https://github.com/rui314) 的 [chibicc](https://github.com/rui314/chibicc) Go 重写，一个五脏俱全的 C 编译器，后端参考 [rvcc](https://github.com/sunshaoce/rvcc) 写了 RISC-V 版本，但中途测试失败没有继续，最终完成的版本只有 x86-64 的内容

[brainfuck](https://github.com/wasuppu/brainfuck): 一个渐进实现的 brainfuck 解释器，包括几个不同版本

[lisgo](https://github.com/wasuppu/lisgo): [norvig lispy](https://norvig.com/lispy.html) 的 Go 重写，文章还有第二篇待做

[chip8](https://github.com/wasuppu/chip8): 参照这篇文章 [Building a CHIP-8 Emulator](https://austinmorlan.com/posts/chip8_emulator/)，指令文档主要参考 [devernay](http://devernay.free.fr/hacks/chip8/) 里的资料，做了一些更改使其通过 [chip8-test-rom](https://github.com/corax89/chip8-test-rom/) 和 [chip8-test-suite](https://github.com/Timendus/chip8-test-suite/) 里的大多数测试。没有实现声音

### 重写的教程

跟随教程，将原语言转写为 Go 实现，没有过多改变

[lc3](https://github.com/wasuppu/lc3): [Write your Own Virtual Machine](https://www.jmeiners.com/lc3-vm/) C -> Go，键盘部分尝试了几种方式，都不完美，最后还是使用了外部库 [keyboard](https://github.com/eiannone/keyboard)

[kilo](https://github.com/wasuppu/kilo): [Build Your Own Text Editor](https://viewsourcecode.org/snaptoken/kilo/) 的 Go 实现，第二章的 raw mode 从 [go/term.git](https://go.googlesource.com/term.git) 和 [go/sys.git](https://go.googlesource.com/sys.git) 拿了代码

[emurv](https://github.com/wasuppu/emurv): 参照 PLCT 的 [KSCO老师带你手搓RISC-V高性能模拟器](https://www.bilibili.com/video/BV1uY4y1D7bJ) 写的，完成到第7课，读取可执行文件部分从之前写的 elf 解析器中拿取了一些代码，指令实现有些问题没有继续向后做

[rvemu](https://github.com/wasuppu/rvemu): 跟着 [Writing a RISC-V Emulator in Rust](https://github.com/d0iasm/book.rvemu) 实现的, 前面的阶段运行正常。最后一步执行 xv6，会运行很长时间然后报错，可能同上也是指令实现问题？我可能需要一个详细的测试

[dnsguide](https://github.com/wasuppu/dnsguide): [dnsguide](https://github.com/EmilHernvall/dnsguide) Rust -> Go，其实还是不懂，照着翻译了一遍，不过代码可以运行

[vulkan-tutorial](https://github.com/wasuppu/vulkan-tutorial): [vulkan-tutorial](https://vulkan-tutorial.com) C++ -> Go, 完成到27课，需要加载模型。在前面 tinyrenderer 项目中编写的 obj 文件加载缺少了对对象 m 的支持，有时间修改下再继续完成

[simpledb](https://github.com/wasuppu/simpledb): [db_tutorial](https://cstack.github.io/db_tutorial/) C -> Go, 用 Go 的测试来替代 Ruby 的 rspec 进行测试用例的编写

[caster](https://github.com/wasuppu/caster): [lodev raycasting](https://lodev.org/cgtutor/raycasting.html) C++ -> Go, 使用 SDL2 来替代作者的 QuickCG，最后一章只做了作者提供代码的增补部分

[smu](https://github.com/wasuppu/smu): [Gottox/smu](https://github.com/Gottox/smu) C -> Golang，支持了 unicode, 本博客使用的 markdown 解析器

[lox](https://github.com/wasuppu/lox): [craftinginterpreters](http://craftinginterpreters.com/) 前半部分的代码实现，不过卡在了第十二章，代码未整理，只有最近的这个版本。跟随 codecrafters 进行了测试，等有时间重新整理各个章节代码并完整实现后再更新

### 浏览一遍的教程

浏览一遍的教程或课程，还是同样的语言，手动敲写或是复制粘贴

[os1000](https://github.com/wasuppu/os1000): [Operating System in 1,000 Lines](https://operating-system-in-1000-lines.vercel.app/) C，没懂，过了一遍

[learn-go-with-tests](https://github.com/wasuppu/learn-go-with-tests): [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests) 刚开始学 Go 的时候第一版没看完，第二版第一部分认真看了，后面变的有些复杂敷衍过去了

[monkey](https://github.com/wasuppu/monkey): [interpreterbook](https://interpreterbook.com/) 只完成了解释器部分，包括第五章，有时间完成编译器部分

[web](https://github.com/wasuppu/web): 看过的 web 开发的课程的汇总，为面试做的准备，不过代码有些混乱，内容也不很完整

[7days](https://github.com/wasuppu/7days): [7days-golang](https://github.com/geektutu/7days-golang) 快速的看了一遍，并没有完全理解，极客兔兔老师写的很好

[gol](https://github.com/wasuppu/gol): [Conway's Game of Life implemented with Go and OpenGL](https://github.com/KyleBanks/conways-gol) 尝试用kilo的经验在终端复刻失败，仓库里是跟作者完全一样的OpenGL版本

[redis-scratch](https://github.com/wasuppu/redis-scratch): [build-redis-from-scratch](https://www.build-redis-from-scratch.dev/) 虽然很短，跟完全的 redis 差的远得多，但代码写的有点启发，为 codecrafters 的 redis 练习找到资料，尽管看完也只做了一点

[gophercises](https://github.com/wasuppu/gophercises): [gophercises](https://gophercises.com/) 在这里了解些奇技淫巧，前面还有跟着敲代码，后面只看视频运行代码了

[learnopengl](https://github.com/wasuppu/learnopengl): [learnopengl-cn's blog](https://learnopengl-cn.github.io/) C++，以前跟写的，光照完成，该到模型加载部分，后面还有挺多

