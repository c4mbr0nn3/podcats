namespace Infrastructure.Persistence;

public class DatabaseSettings
{
    public Provider Provider { get; }
    public IProviderSettings ProviderSettings { get; }

    public DatabaseSettings()
    {
        Provider = GetProvider();
        ProviderSettings = Provider switch
        {
            Provider.Sqlite => new SqliteSettings(),
            Provider.Postgres => new PostgresSettings(),
            _ => throw new Exception("Unsupported database provider")
        };
    }

    private static Provider GetProvider()
    {
        var provider = Environment.GetEnvironmentVariable("DB_PROVIDER") ?? "sqlite";
        return ProviderParser.Parse(provider);
    }
}

public enum Provider
{
    Sqlite = 1,
    Postgres = 2
}

public static class ProviderParser
{
    public static Provider Parse(string provider)
    {
        return provider switch
        {
            "postgres" => Provider.Postgres,
            "sqlite" => Provider.Sqlite,
            _ => throw new Exception("Unsupported database provider")
        };
    }
}

public interface IProviderSettings
{
    public string GetConnectionString();
}

public class SqliteSettings : IProviderSettings
{
    public string GetConnectionString()
    {
        return "Data Source=app.db";
    }
}

public class PostgresSettings : IProviderSettings
{
    public string? Host { get; set; }
    public string? Port { get; set; }
    public string? Database { get; set; }
    public string? Username { get; set; }
    public string? Password { get; set; }

    public string GetConnectionString()
    {
        return $"Host={Host};" +
               $"Port={Port};" +
               $"Database={Database};" +
               $"Username={Username};" +
               $"Password={Password}";
    }
}