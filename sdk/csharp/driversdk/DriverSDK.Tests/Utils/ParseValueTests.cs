using DriverSDK.DmContext.Models;
using DriverSDK.Utils;

namespace DriverSDK.Tests.Utils
{
    [TestFixture]
    public class ParseValueTests
	{
        [Test]
        public void TestDateAndTimeParse()
        {
            string[] dateLayout = { "yyyy-MM-dd", "yyyy.MM.dd", "yyyy/MM/dd" };
            string[] timeLayout = { "HH:mm:ss", "HH-mm-ss", "HH.mm.ss" };

            string res;
            res = ParseValue.ParseDate("2023-07-21", "[yyyy][MM][dd]", dateLayout);
            Assert.That(res, Is.EqualTo("[2023][07][21]"));

            res = ParseValue.ParseDate("2023.07.21", "[yyyy][MM][dd]", dateLayout);
            Assert.That(res, Is.EqualTo("[2023][07][21]"));

            res = ParseValue.ParseDate("2023/07/21", "[yyyy][MM][dd]", dateLayout);
            Assert.That(res, Is.EqualTo("[2023][07][21]"));

            res = ParseValue.ParseDate("18:14:08", "[HH][mm][ss]", timeLayout);
            Assert.That(res, Is.EqualTo("[18][14][08]"));

            res = ParseValue.ParseDate("18-14-08", "[HH][mm][ss]", timeLayout);
            Assert.That(res, Is.EqualTo("[18][14][08]"));

            res = ParseValue.ParseDate("18.14.08", "[HH][mm][ss]", timeLayout);
            Assert.That(res, Is.EqualTo("[18][14][08]"));

            res = ParseValue.ParseDate("2023-07.21", "[yyyy][MM][dd]", dateLayout);
            Assert.That(res, Is.EqualTo(""));

            res = ParseValue.ParseDate("18-14.08", "[HH][mm][ss]", timeLayout);
            Assert.That(res, Is.EqualTo(""));
        }

        [Test]
        public void TestParseArray()
        {
            ArrayType args = new(ParseValue.TYPE_DATE, 2, 5, "(yyyy)[MM]{dd}");
            string[] dates = new string[] { "2023.07.21", "2023.07.22", "2023.07.23" };
            object[]? res = ParseValue.ParseArray(dates, args);

            string[] exps = new string[] { "(2023)[07]{21}", "(2023)[07]{22}", "(2023)[07]{23}" };
            CollectionAssert.AreEqual(exps, res);

            dates = new string[] { "2023.07.21" };
            res = ParseValue.ParseArray(dates, args);
            Assert.That(res, Is.Null);
        }

        [Test]
        public void TestParseEnum()
        {
            EnumType args = new(ParseValue.TYPE_INT, new EnumValue[]
            {
                new EnumValue("e0", "0", "枚举0"),
                new EnumValue("e1", "1", "枚举1"),
                new EnumValue("e2", "2", "枚举2")
            });

            string res = ParseValue.ParseEnum(0, args);
            Assert.That(res, Is.EqualTo("e0"));

            res = ParseValue.ParseEnum(2, args);
            Assert.That(res, Is.EqualTo("e2"));

            res = ParseValue.ParseEnum(2.0, args);
            Assert.That(res, Is.EqualTo(""));

            res = ParseValue.ParseEnum(false, args);
            Assert.That(res, Is.EqualTo(""));
        }

        [Test]
        public void TestParseObject()
        {
            Dictionary<string, ObjectType> args = new()
            {
                { "obj0", new ObjectType("d0", ParseValue.TYPE_INT, "") },
                { "obj1", new ObjectType("d1", ParseValue.TYPE_BOOL, "") },
                { "obj2", new ObjectType("d2", ParseValue.TYPE_FLOAT64, "") },
                { "obj3", new ObjectType("d3", ParseValue.TYPE_DATE, "<yyyy>[MM](dd)") }
            };

            Dictionary<string, object> objs = new()
            {
                { "obj0", 369 },
                { "obj1", false },
                { "obj2", 1.2 },
                { "obj3", "1991-10-01" }
            };

            Dictionary<string, object> exps = new()
            {
                { "obj0", 369 },
                { "obj1", false },
                { "obj2", 1.2 },
                { "obj3", "<1991>[10](01)" }
            };

            Dictionary<string, object>? res = ParseValue.ParseObject(objs, args);

            CollectionAssert.AreEqual(exps, res);

            res = ParseValue.ParseObject(new object(), args);
            Assert.That(res, Is.Null);

            res = ParseValue.ParseObject(objs, new object());
            Assert.That(res, Is.Null);
        }

        [Test]
        public void TestParseBase()
        {
            var t0 = ParseValue.Parse(ParseValue.TYPE_INT,"369","");
            Assert.That(t0, Is.EqualTo(369));

            var t1 = ParseValue.Parse(ParseValue.TYPE_INT32, "369", "");
            Assert.That(t1, Is.EqualTo((Int32)369));

            var t2 = ParseValue.Parse(ParseValue.TYPE_INT16, "369", "");
            Assert.That(t2, Is.EqualTo((Int16)369));

            var t3 = ParseValue.Parse(ParseValue.TYPE_FLOAT32, "369", "");
            Assert.That(t3, Is.EqualTo((float)369));

            var t4 = ParseValue.Parse(ParseValue.TYPE_FLOAT64, "369", "");
            Assert.That(t4, Is.EqualTo((Double)369));

            var t5 = ParseValue.Parse(ParseValue.TYPE_BOOL, "true", "");
            Assert.That(t5, Is.EqualTo(true));

            var t6 = ParseValue.Parse(ParseValue.TYPE_STRING, "369", "");
            Assert.That(t6, Is.EqualTo("369"));
        }
    }
}
