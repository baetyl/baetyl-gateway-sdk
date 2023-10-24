package com.baidu.bce.sdk.utils;

import com.baidu.bce.sdk.context.models.device.ArrayType;
import com.baidu.bce.sdk.context.models.device.EnumType;
import com.baidu.bce.sdk.context.models.device.EnumValue;
import com.baidu.bce.sdk.context.models.device.ObjectType;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.text.ParseException;
import java.util.HashMap;
import java.util.Map;

public class ParseValueTest {
    @Test
    public void testDateAndTimeParse() {
        String[] dateLayout = {"yyyy-MM-dd", "yyyy.MM.dd", "yyyy/MM/dd"};
        String[] timeLayout = {"HH:mm:ss", "HH-mm-ss", "HH.mm.ss"};

        String res;
        res = ParseValue.parseDate("2023-07-21", "[yyyy][MM][dd]", dateLayout);
        Assertions.assertEquals("[2023][07][21]", res);

        res = ParseValue.parseDate("2023.07.21", "[yyyy][MM][dd]", dateLayout);
        Assertions.assertEquals("[2023][07][21]", res);

        res = ParseValue.parseDate("2023/07/21", "[yyyy][MM][dd]", dateLayout);
        Assertions.assertEquals("[2023][07][21]", res);

        res = ParseValue.parseDate("18:14:08", "[HH][mm][ss]", timeLayout);
        Assertions.assertEquals("[18][14][08]", res);

        res = ParseValue.parseDate("18-14-08", "[HH][mm][ss]", timeLayout);
        Assertions.assertEquals("[18][14][08]", res);

        res = ParseValue.parseDate("18.14.08", "[HH][mm][ss]", timeLayout);
        Assertions.assertEquals("[18][14][08]", res);

        res = ParseValue.parseDate("2023-07.21", "[yyyy][MM][dd]", dateLayout);
        Assertions.assertEquals("", res);

        res = ParseValue.parseDate("18-14.08", "[HH][mm][ss]", timeLayout);
        Assertions.assertEquals("", res);
    }

    @Test
    public void testParseArray() throws ParseException {
        ArrayType args = new ArrayType(ParseValue.TYPE_DATE, 2, 5, "(yyyy)[MM]{dd}");
        String[] dates = new String[]{"2023.07.21", "2023.07.22", "2023.07.23"};
        Object[] res = ParseValue.parseArray(dates, args);

        String[] exps = new String[]{"(2023)[07]{21}", "(2023)[07]{22}", "(2023)[07]{23}"};
        Assertions.assertArrayEquals(exps, res);

        dates = new String[]{"2023.07.21"};
        res = ParseValue.parseArray(dates, args);
        Assertions.assertArrayEquals(null, res);
    }

    @Test
    public void testParseEnum() throws ParseException {
        EnumType args = new EnumType(ParseValue.TYPE_INT, new EnumValue[]{
                new EnumValue("e0", "0", "枚举0"),
                new EnumValue("e1", "1", "枚举1"),
                new EnumValue("e2", "2", "枚举2")
        });

        String res = ParseValue.parseEnum(0, args);
        Assertions.assertEquals("e0", res);

        res = ParseValue.parseEnum(2, args);
        Assertions.assertEquals("e2", res);

        res = ParseValue.parseEnum(2.0, args);
        Assertions.assertEquals("", res);

        res = ParseValue.parseEnum(false, args);
        Assertions.assertEquals("", res);
    }

    @Test
    public void testParseObject() throws ParseException {
        Map<String, ObjectType> args = new HashMap<>();
        args.put("obj0", new ObjectType("d0", ParseValue.TYPE_INT, ""));
        args.put("obj1", new ObjectType("d1", ParseValue.TYPE_BOOL, ""));
        args.put("obj2", new ObjectType("d2", ParseValue.TYPE_FLOAT64, ""));
        args.put("obj3", new ObjectType("d3", ParseValue.TYPE_DATE, "<yyyy>[MM](dd)"));

        Map<String, Object> objs = new HashMap<>();
        objs.put("obj0", 369);
        objs.put("obj1", false);
        objs.put("obj2", 1.2);
        objs.put("obj3", "1991-10-01");

        Map<String, Object> exps = new HashMap<>();
        exps.put("obj0", 369);
        exps.put("obj1", false);
        exps.put("obj2", 1.2);
        exps.put("obj3", "<1991>[10](01)");

        Map<String, Object> res = ParseValue.parseObject(objs, args);
        Assertions.assertEquals(exps, res);

        res = ParseValue.parseObject(new Object(), args);
        Assertions.assertNull(res);

        res = ParseValue.parseObject(objs, new Object());
        Assertions.assertNull(res);
    }
}
