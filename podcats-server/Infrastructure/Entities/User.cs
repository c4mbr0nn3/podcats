using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace Infrastructure.Entities;

[Table("users")]
public class User : AuditableEntity
{
    [Column("id")] [Key] public Guid Id { get; set; }
    [Column("first_name")] [Required] public string FirstName { get; set; }
    [Column("last_name")] [Required] public string LastName { get; set; }
    [Column("username")] [Required] public string Username { get; set; }
    [Column("email")] [Required] public string Email { get; set; }
    [Column("role")] [Required] public UserRole Role { get; set; }
    [Column("password_hash")] [Required] public string PasswordHash { get; set; }

    [Column("need_password_change")]
    [Required]
    public bool NeedPasswordChange { get; set; }
}

public enum UserRole
{
    Admin = 1,
    User = 2
}