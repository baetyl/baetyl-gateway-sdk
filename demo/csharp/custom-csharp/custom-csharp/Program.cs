using DriverSDK.Plugin;
using Buildin.Csharp.Collecton;

class Program
{
    static void Main()
    {
        new Serve(new Driver()).Start();
    }
}
