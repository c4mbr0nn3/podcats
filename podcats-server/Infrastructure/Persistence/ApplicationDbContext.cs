using Infrastructure.Entities;
using Microsoft.EntityFrameworkCore;

namespace Infrastructure.Persistence;

public class ApplicationDbContext(
    DbContextOptions<ApplicationDbContext> options,
    DatabaseSettings settings
) : DbContext(options)
{
    public DbSet<User> Users { get; set; }
    public DbSet<Podcast> Podcasts { get; set; }
    public DbSet<Episode> Episodes { get; set; }
    public DbSet<Notification> Notifications { get; set; }

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        base.OnModelCreating(modelBuilder);

        if (settings.Provider == Provider.Postgres) modelBuilder.HasDefaultSchema("public");

        /* modelBuilder.Entity<User>().HasData(new User
        {
            Id = Guid.NewGuid(),
            FirstName = "Super",
            LastName = "User",
            Username = "superuser",
            Email = "example@email.podcats",
            Role = UserRole.Admin,
            PasswordHash = BCrypt.Net.BCrypt.HashPassword("superuser")
        });*/
    }
}