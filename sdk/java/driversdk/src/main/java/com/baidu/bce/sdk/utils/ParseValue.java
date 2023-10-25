package com.baidu.bce.sdk.utils;

import com.baidu.bce.sdk.context.models.device.ArrayType;
import com.baidu.bce.sdk.context.models.device.EnumType;
import com.baidu.bce.sdk.context.models.device.EnumValue;
import com.baidu.bce.sdk.context.models.device.ObjectType;

import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;

public class ParseValue {
    public static final String TYPE_INT = "int";
    public static final String TYPE_INT16 = "int16";
    public static final String TYPE_INT32 = "int32";
    public static final String TYPE_INT64 = "int64";
    public static final String TYPE_FLOAT32 = "float32";
    public static final String TYPE_FLOAT64 = "float64";
    public static final String TYPE_BOOL = "bool";
    public static final String TYPE_STRING = "string";
    public static final String TYPE_TIME = "time";
    public static final String TYPE_DATE = "date";
    public static final String TYPE_ARRAY = "array";
    public static final String TYPE_ENUM = "enum";
    public static final String TYPE_OBJECT = "object";

    private static String[] dateLayout = {
            "yyyy-MM-dd", "yyyy.MM.dd", "yyyy/MM/dd"
    };

    private static String[] timeLayout = {
            "HH:mm:ss", "HH-mm-ss", "HH.mm.ss"
    };

    public static Object parse(String typ, Object value, Object args) throws ParseException {
        try {
            return switch (typ) {
                case TYPE_INT, TYPE_INT32 -> Integer.parseInt(value.toString());
                case TYPE_INT16 -> Short.parseShort(value.toString());
                case TYPE_INT64 -> Long.parseLong(value.toString());
                case TYPE_FLOAT32 -> Float.parseFloat(value.toString());
                case TYPE_FLOAT64 -> Double.parseDouble(value.toString());
                case TYPE_BOOL -> Boolean.parseBoolean(value.toString());
                case TYPE_STRING -> value.toString();
                case TYPE_DATE -> parseDate(value.toString(), args.toString(), dateLayout);
                case TYPE_TIME -> parseDate(value.toString(), args.toString(), timeLayout);
                case TYPE_ARRAY -> parseArray(value, args);
                case TYPE_ENUM -> parseEnum(value, args);
                case TYPE_OBJECT -> parseObject(value, args);
                default -> null;
            };
        } catch (Exception e) {
            throw new ParseException(e.getMessage(), 0);
        }
    }

    public static Map<String, Object> parseObject(Object obj, Object args) throws ParseException {
        Map<String, Object> res = null;
        if (args instanceof Map && obj instanceof Map) {
            res = new HashMap<>();

            for (Map.Entry<?, ?> entry : ((Map<?, ?>) args).entrySet()) {
                if (entry.getKey() instanceof String && entry.getValue() instanceof ObjectType) {
                    Object originVal = ((Map<?, ?>) obj).get(entry.getKey());
                    res.put(((String) entry.getKey()), parse(((ObjectType) entry.getValue()).getType(),
                            originVal, ((ObjectType) entry.getValue()).getFormat()));
                }
            }
        }
        return res;
    }

    public static String parseEnum(Object obj, Object args) throws ParseException {
        if (args instanceof EnumType) {
            for (EnumValue item : ((EnumType) args).getValues()) {
                Object val = parse(((EnumType) args).getType(), item.getValue(), null);
                if (val.equals(obj)) {
                    return item.getName();
                }
            }
        }
        return "";
    }

    public static Object[] parseArray(Object obj, Object args) throws ParseException {
        Object[] res = null;
        if (args instanceof ArrayType) {
            if (obj instanceof Object[]) {
                if (((Object[]) obj).length <= ((ArrayType) args).getMax()
                        && ((Object[]) obj).length >= ((ArrayType) args).getMin()) {
                    res = new Object[((Object[]) obj).length];
                    for (int i = 0; i < ((Object[]) obj).length; i++) {
                        Object t = parse(((ArrayType) args).getType(),
                                ((Object[]) obj)[i], ((ArrayType) args).getFormat());
                        res[i] = t;
                    }
                }
            }
        }
        return res;
    }

    public static String parseDate(String val, String format, String[] layouts) {
        SimpleDateFormat sdf;
        Date date = null;
        for (String layout : layouts) {
            sdf = new SimpleDateFormat(layout);
            try {
                date = sdf.parse(val);
            } catch (ParseException e) {
                continue;
            }
            break;
        }
        if (date != null) {
            sdf = new SimpleDateFormat(format);
            return sdf.format(date);
        }
        return "";
    }

}
