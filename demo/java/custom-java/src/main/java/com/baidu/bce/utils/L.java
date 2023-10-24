package com.baidu.bce.utils;

import java.io.BufferedWriter;
import java.io.FileWriter;
import java.io.IOException;
import java.text.SimpleDateFormat;
import java.util.Date;

public class L {
    private L() {
    }

    private static BufferedWriter bufferedWriter;
    private static SimpleDateFormat sdf;

    static {
        String filePath = "output/custom-java.txt"; // 文件路径
        sdf = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss.SSS");
        try {
            FileWriter fileWriter = new FileWriter(filePath, false);
            bufferedWriter = new BufferedWriter(fileWriter);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public static void info(String s) {
        text("[info] " + s);
    }

    public static void debug(String s) {
        text("[debug] " + s);
    }

    public static void warn(String s) {
        text("[warn] " + s);
    }

    public static void error(String s) {
        text("[error] " + s);
    }

    private static void text(String s) {
        s = "[" + sdf.format(new Date(System.currentTimeMillis())) + "] " + s;
        System.err.println(s);
        if (bufferedWriter != null) {
            try {
                bufferedWriter.write(s);
                bufferedWriter.newLine();
                bufferedWriter.flush();
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
    }
}

