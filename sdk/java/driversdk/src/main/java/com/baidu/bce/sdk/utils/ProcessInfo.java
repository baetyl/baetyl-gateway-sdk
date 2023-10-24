package com.baidu.bce.sdk.utils;

import java.lang.management.ManagementFactory;
import java.lang.management.RuntimeMXBean;

public class ProcessInfo {
    public static long getPidOfProcess() {
        RuntimeMXBean runtimeMXBean = ManagementFactory.getRuntimeMXBean();
        // 获取当前进程的名称
        String processName = runtimeMXBean.getName();
        // 解析出进程 ID
        return Long.parseLong(processName.split("@")[0]);
    }

    public static long getPidOfParentProcess() {
        long processId = getPidOfProcess();
        return getPidOfParentProcessByPid(processId);
    }

    public static long getPidOfParentProcessByPid(long processId) {
        // 根据操作系统的不同，获取父进程 ID 的方式也不同
        String os = System.getProperty("os.name").toLowerCase();

        try {
            if (os.contains("win")) {
                // Windows 系统
                Process process = Runtime.getRuntime().exec(
                        new String[]{String.format("wmic process where ProcessId=%d get ParentProcessId", processId)});
                process.getOutputStream().close();

                java.util.Scanner scanner = new java.util.Scanner(process.getInputStream()).useDelimiter("\\A");
                String output = scanner.hasNext() ? scanner.next() : "";

                // 解析出父进程 ID
                String parentProcessIdStr = output.trim().split("\\s+")[1];
                return Long.parseLong(parentProcessIdStr);
            } else {
                // 非 Windows 系统（Unix/Linux/Mac）
                Process process = Runtime.getRuntime().exec(
                        new String[]{"bash", "-c", "ps -p " + processId + " -o ppid="});
                process.getOutputStream().close();

                java.util.Scanner scanner = new java.util.Scanner(process.getInputStream()).useDelimiter("\\A");
                String output = scanner.hasNext() ? scanner.next() : "";

                // 解析出父进程 ID
                String parentProcessIdStr = output.trim();
                return Long.parseLong(parentProcessIdStr);
            }
        } catch (Exception e) {
            e.printStackTrace();
        }

        return -1; // 获取父进程 ID 失败
    }
}