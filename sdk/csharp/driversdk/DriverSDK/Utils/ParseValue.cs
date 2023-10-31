using DriverSDK.DmContext.Models;
using System.Globalization;

namespace DriverSDK.Utils
{
    public class ParseValue
    {
        public const string TYPE_INT = "int";
        public const string TYPE_INT16 = "int16";
        public const string TYPE_INT32 = "int32";
        public const string TYPE_INT64 = "int64";
        public const string TYPE_FLOAT32 = "float32";
        public const string TYPE_FLOAT64 = "float64";
        public const string TYPE_BOOL = "bool";
        public const string TYPE_STRING = "string";
        public const string TYPE_TIME = "time";
        public const string TYPE_DATE = "date";
        public const string TYPE_ARRAY = "array";
        public const string TYPE_ENUM = "enum";
        public const string TYPE_OBJECT = "object";

        private static readonly string[] DateLayout = { "yyyy-MM-dd", "yyyy.MM.dd", "yyyy/MM/dd" };
        private static readonly string[] TimeLayout = { "HH:mm:ss", "HH-mm-ss", "HH.mm.ss" };

        public static object? Parse(string typ, object value, object args)
        {
            try
            {
                return typ switch
                {
                    TYPE_INT or TYPE_INT32 => Convert.ToInt32(value),
                    TYPE_INT16 => Convert.ToInt16(value),
                    TYPE_INT64 => Convert.ToInt64(value),
                    TYPE_FLOAT32 => Convert.ChangeType(value, typeof(float)),
                    TYPE_FLOAT64 => Convert.ToDouble(value),
                    TYPE_BOOL => Convert.ToBoolean(value),
                    TYPE_STRING => Convert.ToString(value),
                    TYPE_DATE => ParseDate(Convert.ToString(value), Convert.ToString(args), DateLayout),
                    TYPE_TIME => ParseDate(Convert.ToString(value), Convert.ToString(args), TimeLayout),
                    TYPE_ARRAY => ParseArray(value, args),
                    TYPE_ENUM => ParseEnum(value, args),
                    TYPE_OBJECT => ParseObject(value, args),
                    _ => null,
                };
            }
            catch (Exception e)
            {
                throw new InvalidCastException(e.Message);
            }
        }

        public static Dictionary<string, object>? ParseObject(object obj, object args)
        {
            if (args is IDictionary<string, ObjectType> argsDict && obj is Dictionary<string, object> objDict)
            {
                var res = new Dictionary<string, object>();
                foreach (var entry in argsDict)
                {
                    if (entry.Key is string key && entry.Value is ObjectType objectType && objDict.ContainsKey(key))
                    {
                        var originVal = objDict[key];
                        if (objectType.Type != null && originVal != null && objectType.Format != null)
                        {
                            var parRes = Parse(objectType.Type, originVal, objectType.Format);
                            if (parRes != null)
                            {
                                res[key] = parRes;
                            }
                        }
                    }
                }
                return res;
            }
            return null;
        }

        public static string ParseEnum(object obj, object args)
        {
            if (args is EnumType enumType && enumType.Values != null && enumType.Type != null)
            {
                foreach (var item in enumType.Values)
                {
                    var val = Parse(enumType.Type, item.Value, "");
                    if (val != null && val.Equals(obj))
                    {
                        return item.Name;
                    }
                }
            }
            return "";
        }

        public static object[]? ParseArray(object obj, object args)
        {
            if (args is ArrayType arrayType && obj is object[] objArray)
            {
                if (objArray.Length <= arrayType.Max && objArray.Length >= arrayType.Min)
                {
                    var res = new object[objArray.Length];
                    for (int i = 0; i < objArray.Length; i++)
                    {
                        var t = Parse(arrayType.Type, objArray[i], arrayType.Format);
                        if (t != null)
                        {
                            res[i] = t;
                        }
                    }
                    return res;
                }
            }
            return null;
        }

        public static string ParseDate(string? val, string? format, string[] layouts)
        {
            if (val == null || format == null)
            {
                return "";
            }
            foreach (var layout in layouts)
            {
                if (DateTime.TryParseExact(val, layout, CultureInfo.InvariantCulture, DateTimeStyles.None, out DateTime date))
                {
                    return date.ToString(format);
                }
            }
            return "";
        }
    }
}

