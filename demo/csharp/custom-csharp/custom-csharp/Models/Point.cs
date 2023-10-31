namespace Buildin.Csharp.Models
{
    public class Point
    {
        public const int INDEX_TEMPERATURE = 0;
        public const int INDEX_HUMIDITY = 1;
        public const int INDEX_PRESSURE = 2;

        public double Temperature { get; set; } = 0.0;
        public double Humidity { get; set; } = 0.0;
        public double Pressure { get; set; } = 0.0;

        public Point() { }

        public Point(double temperature, double humidity, double pressure)
        {
            Temperature = temperature;
            Humidity = humidity;
            Pressure = pressure;
        }
    }
}

