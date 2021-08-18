package Utils

//
///*
//// C 标志io头文件，你也可以使用里面提供的函数
//#include <stdio.h>
//#include "pycdc/astree.h"
//
//#ifdef WIN32
//#  define PATHSEP '\\'
//#else
//#  define PATHSEP '/'
//#endif
//int PycDcGo(string FileName)
//{
//    PycModule mod;
//    try {
//        mod.loadFromFile(FileName);
//    } catch (std::exception& ex) {
//        fprintf(stderr, "Error loading file %s: %s\n",FileName, ex.what());
//        return 1;
//    }
//    if (!mod.isValid()) {
//        fprintf(stderr, "Could not load file %s\n", FileName);
//        return 1;
//    }
//    const char* dispname = strrchr(FileName, PATHSEP);
//    dispname = (dispname == NULL) ? FileName : dispname + 1;
//    fputs("# Source Generated with Decompyle++\n", pyc_output);
//    fprintf(pyc_output, "# File: %s (Python %d.%d%s)\n\n", dispname, mod.majorVer(), mod.minorVer(),
//            (mod.majorVer() < 3 && mod.isUnicode()) ? " Unicode" : "");
//    try {
//		writepy("file.py");
//        decompyle(mod.code(), &mod);
//    } catch (std::exception& ex) {
//        fprintf(stderr, "Error decompyling %s: %s\n", FileName, ex.what());
//        return 1;
//    }
//
//    return 0;
//}
//
//*/
//import "C"  // 切勿换行再写这个
//
//func Pycdc(FileName string)  {
//	C.PycDcGo(FileName)
//}
