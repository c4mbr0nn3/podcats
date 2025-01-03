using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace Infrastructure.Entities;

public abstract class AuditableEntity
{
    [Column("created_at")] [Required] public DateTime CreatedAt { get; set; } = DateTime.UtcNow;
    [Column("updated_at")] public DateTime UpdatedAt { get; set; }
    [Column("deleted_at")] public DateTime DeletedAt { get; set; }
    [Column("updated_by")] public Guid UpdatedBy { get; set; }
}