namespace DriverSDK.DmContext.Models
{
    public class Message
    {
        public string Kind { get; set; }
        public Dictionary<string, string> Meta { get; set; }
        public object Content { get; set; }

        public Message(string kind, Dictionary<string, string> meta, object content)
        {
            this.Kind = kind;
            this.Meta = meta;
            this.Content = content;
        }
    }
}

