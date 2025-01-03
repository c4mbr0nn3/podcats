using Infrastructure.Entities;
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace Infrastructure.Migrations
{
    /// <inheritdoc />
    public partial class add_super_user : Migration
    {
        private static readonly string[] columns = new[] { "Id", "Email", "FirstName", "LastName", "PasswordHash", "Role", "Username" };

        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.InsertData(
                table: "users",
                columns: columns,
                values: new object[] {

                  Guid.NewGuid(),
                "super@admin.user",
              "Super",
              "User",
               BCrypt.Net.BCrypt.HashPassword("root"),
               UserRole.Admin,
              "superuser"
                 }
            );
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        { }
    }
}
