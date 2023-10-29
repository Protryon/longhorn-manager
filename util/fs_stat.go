package util

const STATFS_ADFS_MAGIC = 0xadf5
const STATFS_AFFS_MAGIC = 0xadff
const STATFS_AFS_MAGIC = 0x5346414F
const STATFS_AUTOFS_MAGIC = 0x0187
const STATFS_BDEVFS_MAGIC = 0x62646576
const STATFS_BEFS_MAGIC = 0x42465331
const STATFS_BFS_MAGIC = 0x1BADFACE
const STATFS_BINFMTFS_MAGIC = 0x42494e4d
const STATFS_BTRFS_MAGIC = 0x9123683E
const STATFS_CEPH_MAGIC = 0x00c36400
const STATFS_CGROUP_MAGIC = 0x27e0eb
const STATFS_CGROUP2_MAGIC = 0x63677270
const STATFS_CIFS_MAGIC = 0xff534d42
const STATFS_CODA_MAGIC = 0x73757245
const STATFS_CONFIGFS_MAGIC = 0x62656570
const STATFS_CRAMFS_MAGIC = 0x28cd3d45
const STATFS_DEBUGFS_MAGIC = 0x64626720
const STATFS_DEVPTS_MAGIC = 0x1cd1
const STATFS_ECRYPTFS_MAGIC = 0xf15f
const STATFS_EFIVARFS_MAGIC = 0xde5e81e4
const STATFS_EFS_MAGIC = 0x414A53
const STATFS_EXOFS_MAGIC = 0x5DF5
const STATFS_EXT2_MAGIC = 0xEF53
const STATFS_EXT3_MAGIC = 0xEF53
const STATFS_EXT4_MAGIC = 0xEF53
const STATFS_F2FS_MAGIC = 0xF2F52010
const STATFS_FUSE_MAGIC = 0x65735546
const STATFS_FUTEXFS_MAGIC = 0xBAD1DEA
const STATFS_GFS2_MAGIC = 0x01161970
const STATFS_HFSPLUS_MAGIC = 0x482b
const STATFS_HOSTFS_MAGIC = 0x00c0ffee
const STATFS_HPFS_MAGIC = 0xf995e849
const STATFS_HPPFS_MAGIC = 0xb00000ee
const STATFS_HUGETLBFS_MAGIC = 0x958458f6
const STATFS_ISOFS_MAGIC = 0x9660
const STATFS_JFFS2_MAGIC = 0x72b6
const STATFS_JFS_MAGIC = 0x3153464a
const STATFS_LOGFS_MAGIC = 0xc97e8168
const STATFS_MINIX2_MAGIC = 0x2468
const STATFS_MINIX2_MAGIC2 = 0x2478
const STATFS_MINIX3_MAGIC = 0x4d5a
const STATFS_MINIX_MAGIC = 0x137F
const STATFS_MINIX_MAGIC2 = 0x138F
const STATFS_MQUEUE_MAGIC = 0x19800202
const STATFS_MSDOS_MAGIC = 0x4d44
const STATFS_NCP_MAGIC = 0x564c
const STATFS_NFS_MAGIC = 0x6969
const STATFS_NILFS_MAGIC = 0x3434
const STATFS_NTFS_MAGIC = 0x5346544e
const STATFS_OCFS2_MAGIC = 0x7461636f
const STATFS_OMFS_MAGIC = 0xC2993D87
const STATFS_OPENPROMFS_MAGIC = 0x9fa1
const STATFS_PIPEFS_MAGIC = 0x50495045
const STATFS_PROC_MAGIC = 0x9fa0
const STATFS_PSTOREFS_MAGIC = 0x6165676C
const STATFS_QNX4_MAGIC = 0x002f
const STATFS_QNX6_MAGIC = 0x68191122
const STATFS_RAMFS_MAGIC = 0x858458f6
const STATFS_REISERFS_MAGIC = 0x52654973
const STATFS_ROMFS_MAGIC = 0x7275
const STATFS_SECURITYFS_MAGIC = 0x73636673
const STATFS_SELINUXFS_MAGIC = 0xf97cff8c
const STATFS_SMACKFS_MAGIC = 0x43415d53
const STATFS_SMB_MAGIC = 0x517B
const STATFS_SOCKFS_MAGIC = 0x534F434B
const STATFS_SQUASHFS_MAGIC = 0x73717368
const STATFS_SYSFS_MAGIC = 0x62656572
const STATFS_TMPFS_MAGIC = 0x01021994
const STATFS_UBIFS_MAGIC = 0x24051905
const STATFS_UDF_MAGIC = 0x15013346
const STATFS_UFS2_MAGIC = 0x19540119
const STATFS_UFS_MAGIC = 0x00011954
const STATFS_V9FS_MAGIC = 0x01021997
const STATFS_VXFS_MAGIC = 0xa501FCF5
const STATFS_XENFS_MAGIC = 0xabba1974
const STATFS_XFS_MAGIC = 0x58465342

func FsTypeName(x int64) string {
	if x == STATFS_ADFS_MAGIC {
		return "adfs"
	} else if x == STATFS_AFFS_MAGIC {
		return "affs"
	} else if x == STATFS_AFS_MAGIC {
		return "afs"
	} else if x == STATFS_AUTOFS_MAGIC {
		return "autofs"
	} else if x == STATFS_BDEVFS_MAGIC {
		return "bdev"
	} else if x == STATFS_BEFS_MAGIC {
		return "befs"
	} else if x == STATFS_BFS_MAGIC {
		return "befs"
	} else if x == STATFS_BINFMTFS_MAGIC {
		return "binfmt_misc"
	} else if x == STATFS_BTRFS_MAGIC {
		return "btrfs"
	} else if x == STATFS_CEPH_MAGIC {
		return "ceph"
	} else if x == STATFS_CGROUP_MAGIC {
		return "cgroup"
	} else if x == STATFS_CIFS_MAGIC {
		return "cifs"
	} else if x == STATFS_CODA_MAGIC {
		return "coda"
	} else if x == STATFS_CONFIGFS_MAGIC {
		return "configfs"
	} else if x == STATFS_CRAMFS_MAGIC {
		return "cramfs"
	} else if x == STATFS_DEBUGFS_MAGIC {
		return "debugfs"
	} else if x == STATFS_DEVPTS_MAGIC {
		return "devpts"
	} else if x == STATFS_ECRYPTFS_MAGIC {
		return "ecryptfs"
	} else if x == STATFS_EFIVARFS_MAGIC {
		return "efivarfs"
	} else if x == STATFS_EFS_MAGIC {
		return "efs"
	} else if x == STATFS_EXOFS_MAGIC {
		return "exofs"
	} else if x == STATFS_EXT4_MAGIC {
		return "ext4"
	} else if x == STATFS_F2FS_MAGIC {
		return "f2fs"
	} else if x == STATFS_FUSE_MAGIC {
		return "fuse"
	} else if x == STATFS_FUTEXFS_MAGIC {
		return "futexfs"
	} else if x == STATFS_GFS2_MAGIC {
		return "gfs2"
	} else if x == STATFS_HFSPLUS_MAGIC {
		return "hfsplus"
	} else if x == STATFS_HOSTFS_MAGIC {
		return "hostfs"
	} else if x == STATFS_HPFS_MAGIC {
		return "hpfs"
	} else if x == STATFS_HPPFS_MAGIC {
		return "hppfs"
	} else if x == STATFS_HUGETLBFS_MAGIC {
		return "hugetlbfs"
	} else if x == STATFS_ISOFS_MAGIC {
		return "iso9660"
	} else if x == STATFS_JFFS2_MAGIC {
		return "jffs2"
	} else if x == STATFS_JFS_MAGIC {
		return "jfs"
	} else if x == STATFS_LOGFS_MAGIC {
		return "logfs"
	} else if x == STATFS_MINIX2_MAGIC || x == STATFS_MINIX2_MAGIC2 || x == STATFS_MINIX3_MAGIC || x == STATFS_MINIX_MAGIC || x == STATFS_MINIX_MAGIC2 {
		return "minix"
	} else if x == STATFS_MQUEUE_MAGIC {
		return "mqueue"
	} else if x == STATFS_MSDOS_MAGIC {
		return "vfat"
	} else if x == STATFS_NCP_MAGIC {
		return "ncp"
	} else if x == STATFS_NFS_MAGIC {
		return "nfs"
	} else if x == STATFS_NILFS_MAGIC {
		return "nilfs2"
	} else if x == STATFS_NTFS_MAGIC {
		return "ntfs"
	} else if x == STATFS_OCFS2_MAGIC {
		return "ocfs2"
	} else if x == STATFS_OMFS_MAGIC {
		return "omfs"
	} else if x == STATFS_OPENPROMFS_MAGIC {
		return "openpromfs"
	} else if x == STATFS_PIPEFS_MAGIC {
		return "pipefs"
	} else if x == STATFS_PROC_MAGIC {
		return "proc"
	} else if x == STATFS_PSTOREFS_MAGIC {
		return "pstore"
	} else if x == STATFS_QNX4_MAGIC {
		return "qnx4"
	} else if x == STATFS_QNX6_MAGIC {
		return "qnx6"
	} else if x == STATFS_RAMFS_MAGIC {
		return "ramfs"
	} else if x == STATFS_REISERFS_MAGIC {
		return "reiser4"
	} else if x == STATFS_ROMFS_MAGIC {
		return "romfs"
	} else if x == STATFS_SECURITYFS_MAGIC {
		return "securityfs"
	} else if x == STATFS_SELINUXFS_MAGIC {
		return "selinuxfs"
	} else if x == STATFS_SMACKFS_MAGIC {
		return "smackfs"
	} else if x == STATFS_SMB_MAGIC {
		return "smb"
	} else if x == STATFS_SOCKFS_MAGIC {
		return "sockfs"
	} else if x == STATFS_SQUASHFS_MAGIC {
		return "squashfs"
	} else if x == STATFS_SYSFS_MAGIC {
		return "sysfs"
	} else if x == STATFS_TMPFS_MAGIC {
		return "tmpfs"
	} else if x == STATFS_UBIFS_MAGIC {
		return "ubifs"
	} else if x == STATFS_UDF_MAGIC {
		return "udf"
	} else if x == STATFS_UFS_MAGIC || x == STATFS_UFS2_MAGIC {
		return "ufs"
	} else if x == STATFS_V9FS_MAGIC {
		return "9p"
	} else if x == STATFS_VXFS_MAGIC {
		return "vxfs"
	} else if x == STATFS_XENFS_MAGIC {
		return "xenfs"
	} else if x == STATFS_XFS_MAGIC {
		return "xfs"
	}
	return "unknown"
}
