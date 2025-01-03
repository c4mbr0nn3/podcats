using Infrastructure.Persistence;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.DependencyInjection;

namespace Infrastructure;

public static class DependencyInjection
{
    public static IServiceCollection AddInfrastructure(this IServiceCollection services)
    {
        var dbSettings = new DatabaseSettings();

        services.AddSingleton(dbSettings);

        var provider = dbSettings.Provider;
        var connectionString = dbSettings.ProviderSettings.GetConnectionString();
        services.AddDbContext<ApplicationDbContext>((_, o) =>
        {
            if (provider == Provider.Postgres) o.UseNpgsql(connectionString);
            else o.UseSqlite(connectionString);
        });

        return services;
    }
}