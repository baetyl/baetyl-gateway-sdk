using Buildin.Csharp.Models;
using DriverSDK.Utils;

namespace Buildin.Csharp.Collecton
{
    public class Simulator
    {
        private const int SIMULATOR_INTERVAL = 10;
        private string name;
        private Point point;
        private Random random;

        public Simulator(string name)
        {
            this.random = new Random(Environment.TickCount);

            this.name = name;
            this.point = new Point();
            this.point.Temperature = this.random.Next(100);
            this.point.Humidity = this.random.Next(100);
            this.point.Pressure = this.random.Next(100);

            GenerateSimulateData();
        }

        public void Set(int index, double val)
        {
            switch (index)
            {
                case Point.INDEX_TEMPERATURE:
                    this.point.Temperature = val;
                    break;
                case Point.INDEX_HUMIDITY:
                    this.point.Humidity = val;
                    break;
                case Point.INDEX_PRESSURE:
                    this.point.Pressure = val;
                    break;
            }
        }

        public object Get(int index)
        {
            switch (index)
            {
                case Point.INDEX_TEMPERATURE:
                    return this.point.Temperature;
                case Point.INDEX_HUMIDITY:
                    return this.point.Humidity;
                case Point.INDEX_PRESSURE:
                    return this.point.Pressure;
            }
            throw new Exception("unsupported point information");
        }

        private void GenerateSimulateData()
        {
            new Thread(() =>
            {
                while (true)
                {
                    this.point.Temperature = this.random.Next(100);
                    this.point.Humidity = this.random.Next(100);
                    this.point.Pressure = this.random.Next(100);

                    L.Debug("Simulate generate random temperature = " + this.point.Temperature);
                    L.Debug("Simulate generate random humidity = " + this.point.Humidity);
                    L.Debug("Simulate generate random pressure = " + this.point.Pressure);

                    Thread.Sleep(TimeSpan.FromSeconds(SIMULATOR_INTERVAL));
                }
            }).Start();
        }
    }
}

