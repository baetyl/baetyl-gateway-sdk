using System.Diagnostics;

namespace DriverSDK.Utils
{
    public class L
    {
        private static readonly string dateFormat = "yyyy-MM-dd HH:mm:ss.fff";

        public static void Info(string s)
        {
            Text("[info] " + s);
        }

        public static void Debug(string s)
        {
            Text("[debug] " + s);
        }

        public static void Warn(string s)
        {
            WithTrace("[warn] " + s);
        }

        public static void Error(string s)
        {
            WithTrace("[error] " + s);
        }

        private static void Text(string s)
        {
            s = "[" + DateTime.Now.ToString(dateFormat) + "] " + s;
            Console.Error.WriteLine(s);
        }

        private static void WithTrace(string s)
        {
            StackTrace stackTrace = new StackTrace();
            StackFrame[] stackFrames = stackTrace.GetFrames();

            var trace = " ( ";
            foreach (StackFrame frame in stackFrames)
            {
                var method = frame.GetMethod();
                var methodName = "";
                if (method != null)
                {
                    methodName = method.Name;
                }

                var fileName = frame.GetFileName();
                var lineNumber = frame.GetFileLineNumber();

                trace += string.Format("{0}:{1}:{2}  ", fileName, lineNumber, methodName);
            }
            Text(s + trace + ")");
        }
    }
}

