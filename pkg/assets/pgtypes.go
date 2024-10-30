package assets

// PGType represents a PostgreSQL type with a type name and an OID.
type PGType struct {
	Typname string
	OID     int
}

// PGTypes is a slice of PGType representing constant PostgreSQL types.
var PGTypes = []PGType{
	{
		Typname: "bool",
		OID:     16,
	},
	{
		Typname: "bytea",
		OID:     17,
	},
	{
		Typname: "char",
		OID:     18,
	},
	{
		Typname: "name",
		OID:     19,
	},
	{
		Typname: "int8",
		OID:     20,
	},
	{
		Typname: "int2",
		OID:     21,
	},
	{
		Typname: "int2vector",
		OID:     22,
	},
	{
		Typname: "int4",
		OID:     23,
	},
	{
		Typname: "regproc",
		OID:     24,
	},
	{
		Typname: "text",
		OID:     25,
	},
	{
		Typname: "OID",
		OID:     26,
	},
	{
		Typname: "tid",
		OID:     27,
	},
	{
		Typname: "xid",
		OID:     28,
	},
	{
		Typname: "cid",
		OID:     29,
	},
	{
		Typname: "OIDvector",
		OID:     30,
	},
	{
		Typname: "pg_ddl_command",
		OID:     32,
	},
	{
		Typname: "pg_type",
		OID:     71,
	},
	{
		Typname: "pg_attribute",
		OID:     75,
	},
	{
		Typname: "pg_proc",
		OID:     81,
	},
	{
		Typname: "pg_class",
		OID:     83,
	},
	{
		Typname: "json",
		OID:     114,
	},
	{
		Typname: "xml",
		OID:     142,
	},
	{
		Typname: "_xml",
		OID:     143,
	},
	{
		Typname: "pg_node_tree",
		OID:     194,
	},
	{
		Typname: "_json",
		OID:     199,
	},
	{
		Typname: "_pg_type",
		OID:     210,
	},
	{
		Typname: "table_am_handler",
		OID:     269,
	},
	{
		Typname: "_pg_attribute",
		OID:     270,
	},
	{
		Typname: "_xid8",
		OID:     271,
	},
	{
		Typname: "_pg_proc",
		OID:     272,
	},
	{
		Typname: "_pg_class",
		OID:     273,
	},
	{
		Typname: "index_am_handler",
		OID:     325,
	},
	{
		Typname: "point",
		OID:     600,
	},
	{
		Typname: "lseg",
		OID:     601,
	},
	{
		Typname: "path",
		OID:     602,
	},
	{
		Typname: "box",
		OID:     603,
	},
	{
		Typname: "polygon",
		OID:     604,
	},
	{
		Typname: "line",
		OID:     628,
	},
	{
		Typname: "_line",
		OID:     629,
	},
	{
		Typname: "cidr",
		OID:     650,
	},
	{
		Typname: "_cidr",
		OID:     651,
	},
	{
		Typname: "float4",
		OID:     700,
	},
	{
		Typname: "float8",
		OID:     701,
	},
	{
		Typname: "unknown",
		OID:     705,
	},
	{
		Typname: "circle",
		OID:     718,
	},
	{
		Typname: "_circle",
		OID:     719,
	},
	{
		Typname: "macaddr8",
		OID:     774,
	},
	{
		Typname: "_macaddr8",
		OID:     775,
	},
	{
		Typname: "money",
		OID:     790,
	},
	{
		Typname: "_money",
		OID:     791,
	},
	{
		Typname: "macaddr",
		OID:     829,
	},
	{
		Typname: "inet",
		OID:     869,
	},
	{
		Typname: "_bool",
		OID:     1000,
	},
	{
		Typname: "_bytea",
		OID:     1001,
	},
	{
		Typname: "_char",
		OID:     1002,
	},
	{
		Typname: "_name",
		OID:     1003,
	},
	{
		Typname: "_int2",
		OID:     1005,
	},
	{
		Typname: "_int2vector",
		OID:     1006,
	},
	{
		Typname: "_int4",
		OID:     1007,
	},
	{
		Typname: "_regproc",
		OID:     1008,
	},
	{
		Typname: "_text",
		OID:     1009,
	},
	{
		Typname: "_tid",
		OID:     1010,
	},
	{
		Typname: "_xid",
		OID:     1011,
	},
	{
		Typname: "_cid",
		OID:     1012,
	},
	{
		Typname: "_OIDvector",
		OID:     1013,
	},
	{
		Typname: "_bpchar",
		OID:     1014,
	},
	{
		Typname: "_varchar",
		OID:     1015,
	},
	{
		Typname: "_int8",
		OID:     1016,
	},
	{
		Typname: "_point",
		OID:     1017,
	},
	{
		Typname: "_lseg",
		OID:     1018,
	},
	{
		Typname: "_path",
		OID:     1019,
	},
	{
		Typname: "_box",
		OID:     1020,
	},
	{
		Typname: "_float4",
		OID:     1021,
	},
	{
		Typname: "_float8",
		OID:     1022,
	},
	{
		Typname: "_polygon",
		OID:     1027,
	},
	{
		Typname: "_OID",
		OID:     1028,
	},
	{
		Typname: "aclitem",
		OID:     1033,
	},
	{
		Typname: "_aclitem",
		OID:     1034,
	},
	{
		Typname: "_macaddr",
		OID:     1040,
	},
	{
		Typname: "_inet",
		OID:     1041,
	},
	{
		Typname: "bpchar",
		OID:     1042,
	},
	{
		Typname: "varchar",
		OID:     1043,
	},
	{
		Typname: "date",
		OID:     1082,
	},
	{
		Typname: "time",
		OID:     1083,
	},
	{
		Typname: "timestamp",
		OID:     1114,
	},
	{
		Typname: "_timestamp",
		OID:     1115,
	},
	{
		Typname: "_date",
		OID:     1182,
	},
	{
		Typname: "_time",
		OID:     1183,
	},
	{
		Typname: "timestamptz",
		OID:     1184,
	},
	{
		Typname: "_timestamptz",
		OID:     1185,
	},
	{
		Typname: "interval",
		OID:     1186,
	},
	{
		Typname: "_interval",
		OID:     1187,
	},
	{
		Typname: "_numeric",
		OID:     1231,
	},
	{
		Typname: "pg_database",
		OID:     1248,
	},
	{
		Typname: "_cstring",
		OID:     1263,
	},
	{
		Typname: "timetz",
		OID:     1266,
	},
	{
		Typname: "_timetz",
		OID:     1270,
	},
	{
		Typname: "bit",
		OID:     1560,
	},
	{
		Typname: "_bit",
		OID:     1561,
	},
	{
		Typname: "varbit",
		OID:     1562,
	},
	{
		Typname: "_varbit",
		OID:     1563,
	},
	{
		Typname: "numeric",
		OID:     1700,
	},
	{
		Typname: "refcursor",
		OID:     1790,
	},
	{
		Typname: "_refcursor",
		OID:     2201,
	},
	{
		Typname: "regprocedure",
		OID:     2202,
	},
	{
		Typname: "regoper",
		OID:     2203,
	},
	{
		Typname: "regoperator",
		OID:     2204,
	},
	{
		Typname: "regclass",
		OID:     2205,
	},
	{
		Typname: "regtype",
		OID:     2206,
	},
	{
		Typname: "_regprocedure",
		OID:     2207,
	},
	{
		Typname: "_regoper",
		OID:     2208,
	},
	{
		Typname: "_regoperator",
		OID:     2209,
	},
	{
		Typname: "_regclass",
		OID:     2210,
	},
	{
		Typname: "_regtype",
		OID:     2211,
	},
	{
		Typname: "record",
		OID:     2249,
	},
	{
		Typname: "cstring",
		OID:     2275,
	},
	{
		Typname: "any",
		OID:     2276,
	},
	{
		Typname: "anyarray",
		OID:     2277,
	},
	{
		Typname: "vOID",
		OID:     2278,
	},
	{
		Typname: "trigger",
		OID:     2279,
	},
	{
		Typname: "language_handler",
		OID:     2280,
	},
	{
		Typname: "internal",
		OID:     2281,
	},
	{
		Typname: "anyelement",
		OID:     2283,
	},
	{
		Typname: "_record",
		OID:     2287,
	},
	{
		Typname: "anynonarray",
		OID:     2776,
	},
	{
		Typname: "pg_authid",
		OID:     2842,
	},
	{
		Typname: "pg_auth_members",
		OID:     2843,
	},
	{
		Typname: "_txid_snapshot",
		OID:     2949,
	},
	{
		Typname: "uuid",
		OID:     2950,
	},
	{
		Typname: "_uuid",
		OID:     2951,
	},
	{
		Typname: "txid_snapshot",
		OID:     2970,
	},
	{
		Typname: "fdw_handler",
		OID:     3115,
	},
	{
		Typname: "pg_lsn",
		OID:     3220,
	},
	{
		Typname: "_pg_lsn",
		OID:     3221,
	},
	{
		Typname: "tsm_handler",
		OID:     3310,
	},
	{
		Typname: "pg_ndistinct",
		OID:     3361,
	},
	{
		Typname: "pg_dependencies",
		OID:     3402,
	},
	{
		Typname: "anyenum",
		OID:     3500,
	},
	{
		Typname: "tsvector",
		OID:     3614,
	},
	{
		Typname: "tsquery",
		OID:     3615,
	},
	{
		Typname: "gtsvector",
		OID:     3642,
	},
	{
		Typname: "_tsvector",
		OID:     3643,
	},
	{
		Typname: "_gtsvector",
		OID:     3644,
	},
	{
		Typname: "_tsquery",
		OID:     3645,
	},
	{
		Typname: "regconfig",
		OID:     3734,
	},
	{
		Typname: "_regconfig",
		OID:     3735,
	},
	{
		Typname: "regdictionary",
		OID:     3769,
	},
	{
		Typname: "_regdictionary",
		OID:     3770,
	},
	{
		Typname: "jsonb",
		OID:     3802,
	},
	{
		Typname: "_jsonb",
		OID:     3807,
	},
	{
		Typname: "anyrange",
		OID:     3831,
	},
	{
		Typname: "event_trigger",
		OID:     3838,
	},
	{
		Typname: "int4range",
		OID:     3904,
	},
	{
		Typname: "_int4range",
		OID:     3905,
	},
	{
		Typname: "numrange",
		OID:     3906,
	},
	{
		Typname: "_numrange",
		OID:     3907,
	},
	{
		Typname: "tsrange",
		OID:     3908,
	},
	{
		Typname: "_tsrange",
		OID:     3909,
	},
	{
		Typname: "tstzrange",
		OID:     3910,
	},
	{
		Typname: "_tstzrange",
		OID:     3911,
	},
	{
		Typname: "daterange",
		OID:     3912,
	},
	{
		Typname: "_daterange",
		OID:     3913,
	},
	{
		Typname: "int8range",
		OID:     3926,
	},
	{
		Typname: "_int8range",
		OID:     3927,
	},
	{
		Typname: "pg_shseclabel",
		OID:     4066,
	},
	{
		Typname: "jsonpath",
		OID:     4072,
	},
	{
		Typname: "_jsonpath",
		OID:     4073,
	},
	{
		Typname: "regnamespace",
		OID:     4089,
	},
	{
		Typname: "_regnamespace",
		OID:     4090,
	},
	{
		Typname: "regrole",
		OID:     4096,
	},
	{
		Typname: "_regrole",
		OID:     4097,
	},
	{
		Typname: "regcollation",
		OID:     4191,
	},
	{
		Typname: "_regcollation",
		OID:     4192,
	},
	{
		Typname: "int4multirange",
		OID:     4451,
	},
	{
		Typname: "nummultirange",
		OID:     4532,
	},
	{
		Typname: "tsmultirange",
		OID:     4533,
	},
	{
		Typname: "tstzmultirange",
		OID:     4534,
	},
	{
		Typname: "datemultirange",
		OID:     4535,
	},
	{
		Typname: "int8multirange",
		OID:     4536,
	},
	{
		Typname: "anymultirange",
		OID:     4537,
	},
	{
		Typname: "anycompatiblemultirange",
		OID:     4538,
	},
	{
		Typname: "pg_brin_bloom_summary",
		OID:     4600,
	},
	{
		Typname: "pg_brin_minmax_multi_summary",
		OID:     4601,
	},
	{
		Typname: "pg_mcv_list",
		OID:     5017,
	},
	{
		Typname: "pg_snapshot",
		OID:     5038,
	},
	{
		Typname: "_pg_snapshot",
		OID:     5039,
	},
	{
		Typname: "xid8",
		OID:     5069,
	},
	{
		Typname: "anycompatible",
		OID:     5077,
	},
	{
		Typname: "anycompatiblearray",
		OID:     5078,
	},
	{
		Typname: "anycompatiblenonarray",
		OID:     5079,
	},
	{
		Typname: "anycompatiblerange",
		OID:     5080,
	},
	{
		Typname: "pg_subscription",
		OID:     6101,
	},
	{
		Typname: "_int4multirange",
		OID:     6150,
	},
	{
		Typname: "_nummultirange",
		OID:     6151,
	},
	{
		Typname: "_tsmultirange",
		OID:     6152,
	},
	{
		Typname: "_tstzmultirange",
		OID:     6153,
	},
	{
		Typname: "_datemultirange",
		OID:     6155,
	},
	{
		Typname: "_int8multirange",
		OID:     6157,
	},
	{
		Typname: "_pg_attrdef",
		OID:     10000,
	},
	{
		Typname: "pg_attrdef",
		OID:     10001,
	},
	{
		Typname: "_pg_constraint",
		OID:     10002,
	},
	{
		Typname: "pg_constraint",
		OID:     10003,
	},
	{
		Typname: "_pg_inherits",
		OID:     10004,
	},
	{
		Typname: "pg_inherits",
		OID:     10005,
	},
	{
		Typname: "_pg_index",
		OID:     10006,
	},
	{
		Typname: "pg_index",
		OID:     10007,
	},
	{
		Typname: "_pg_operator",
		OID:     10008,
	},
	{
		Typname: "pg_operator",
		OID:     10009,
	},
	{
		Typname: "_pg_opfamily",
		OID:     10010,
	},
	{
		Typname: "pg_opfamily",
		OID:     10011,
	},
	{
		Typname: "_pg_opclass",
		OID:     10012,
	},
	{
		Typname: "pg_opclass",
		OID:     10013,
	},
	{
		Typname: "_pg_am",
		OID:     10014,
	},
	{
		Typname: "pg_am",
		OID:     10015,
	},
	{
		Typname: "_pg_amop",
		OID:     10016,
	},
	{
		Typname: "pg_amop",
		OID:     10017,
	},
	{
		Typname: "_pg_amproc",
		OID:     10018,
	},
	{
		Typname: "pg_amproc",
		OID:     10019,
	},
	{
		Typname: "_pg_language",
		OID:     10020,
	},
	{
		Typname: "pg_language",
		OID:     10021,
	},
	{
		Typname: "_pg_largeobject_metadata",
		OID:     10022,
	},
	{
		Typname: "pg_largeobject_metadata",
		OID:     10023,
	},
	{
		Typname: "_pg_largeobject",
		OID:     10024,
	},
	{
		Typname: "pg_largeobject",
		OID:     10025,
	},
	{
		Typname: "_pg_aggregate",
		OID:     10026,
	},
	{
		Typname: "pg_aggregate",
		OID:     10027,
	},
	{
		Typname: "_pg_statistic",
		OID:     10028,
	},
	{
		Typname: "pg_statistic",
		OID:     10029,
	},
	{
		Typname: "_pg_statistic_ext",
		OID:     10030,
	},
	{
		Typname: "pg_statistic_ext",
		OID:     10031,
	},
	{
		Typname: "_pg_statistic_ext_data",
		OID:     10032,
	},
	{
		Typname: "pg_statistic_ext_data",
		OID:     10033,
	},
	{
		Typname: "_pg_rewrite",
		OID:     10034,
	},
	{
		Typname: "pg_rewrite",
		OID:     10035,
	},
	{
		Typname: "_pg_trigger",
		OID:     10036,
	},
	{
		Typname: "pg_trigger",
		OID:     10037,
	},
	{
		Typname: "_pg_event_trigger",
		OID:     10038,
	},
	{
		Typname: "pg_event_trigger",
		OID:     10039,
	},
	{
		Typname: "_pg_description",
		OID:     10040,
	},
	{
		Typname: "pg_description",
		OID:     10041,
	},
	{
		Typname: "_pg_cast",
		OID:     10042,
	},
	{
		Typname: "pg_cast",
		OID:     10043,
	},
	{
		Typname: "_pg_enum",
		OID:     10044,
	},
	{
		Typname: "pg_enum",
		OID:     10045,
	},
	{
		Typname: "_pg_namespace",
		OID:     10046,
	},
	{
		Typname: "pg_namespace",
		OID:     10047,
	},
	{
		Typname: "_pg_conversion",
		OID:     10048,
	},
	{
		Typname: "pg_conversion",
		OID:     10049,
	},
	{
		Typname: "_pg_depend",
		OID:     10050,
	},
	{
		Typname: "pg_depend",
		OID:     10051,
	},
	{
		Typname: "_pg_database",
		OID:     10052,
	},
	{
		Typname: "_pg_db_role_setting",
		OID:     10053,
	},
	{
		Typname: "pg_db_role_setting",
		OID:     10054,
	},
	{
		Typname: "_pg_tablespace",
		OID:     10055,
	},
	{
		Typname: "pg_tablespace",
		OID:     10056,
	},
	{
		Typname: "_pg_authid",
		OID:     10057,
	},
	{
		Typname: "_pg_auth_members",
		OID:     10058,
	},
	{
		Typname: "_pg_shdepend",
		OID:     10059,
	},
	{
		Typname: "pg_shdepend",
		OID:     10060,
	},
	{
		Typname: "_pg_shdescription",
		OID:     10061,
	},
	{
		Typname: "pg_shdescription",
		OID:     10062,
	},
	{
		Typname: "_pg_ts_config",
		OID:     10063,
	},
	{
		Typname: "pg_ts_config",
		OID:     10064,
	},
	{
		Typname: "_pg_ts_config_map",
		OID:     10065,
	},
	{
		Typname: "pg_ts_config_map",
		OID:     10066,
	},
	{
		Typname: "_pg_ts_dict",
		OID:     10067,
	},
	{
		Typname: "pg_ts_dict",
		OID:     10068,
	},
	{
		Typname: "_pg_ts_parser",
		OID:     10069,
	},
	{
		Typname: "pg_ts_parser",
		OID:     10070,
	},
	{
		Typname: "_pg_ts_template",
		OID:     10071,
	},
	{
		Typname: "pg_ts_template",
		OID:     10072,
	},
	{
		Typname: "_pg_extension",
		OID:     10073,
	},
	{
		Typname: "pg_extension",
		OID:     10074,
	},
	{
		Typname: "_pg_foreign_data_wrapper",
		OID:     10075,
	},
	{
		Typname: "pg_foreign_data_wrapper",
		OID:     10076,
	},
	{
		Typname: "_pg_foreign_server",
		OID:     10077,
	},
	{
		Typname: "pg_foreign_server",
		OID:     10078,
	},
	{
		Typname: "_pg_user_mapping",
		OID:     10079,
	},
	{
		Typname: "pg_user_mapping",
		OID:     10080,
	},
	{
		Typname: "_pg_foreign_table",
		OID:     10081,
	},
	{
		Typname: "pg_foreign_table",
		OID:     10082,
	},
	{
		Typname: "_pg_policy",
		OID:     10083,
	},
	{
		Typname: "pg_policy",
		OID:     10084,
	},
	{
		Typname: "_pg_replication_origin",
		OID:     10085,
	},
	{
		Typname: "pg_replication_origin",
		OID:     10086,
	},
	{
		Typname: "_pg_default_acl",
		OID:     10087,
	},
	{
		Typname: "pg_default_acl",
		OID:     10088,
	},
	{
		Typname: "_pg_init_privs",
		OID:     10089,
	},
	{
		Typname: "pg_init_privs",
		OID:     10090,
	},
	{
		Typname: "_pg_seclabel",
		OID:     10091,
	},
	{
		Typname: "pg_seclabel",
		OID:     10092,
	},
	{
		Typname: "_pg_shseclabel",
		OID:     10093,
	},
	{
		Typname: "_pg_collation",
		OID:     10094,
	},
	{
		Typname: "pg_collation",
		OID:     10095,
	},
	{
		Typname: "_pg_parameter_acl",
		OID:     10096,
	},
	{
		Typname: "pg_parameter_acl",
		OID:     10097,
	},
	{
		Typname: "_pg_partitioned_table",
		OID:     10098,
	},
	{
		Typname: "pg_partitioned_table",
		OID:     10099,
	},
	{
		Typname: "_pg_range",
		OID:     10100,
	},
	{
		Typname: "pg_range",
		OID:     10101,
	},
	{
		Typname: "_pg_transform",
		OID:     10102,
	},
	{
		Typname: "pg_transform",
		OID:     10103,
	},
	{
		Typname: "_pg_sequence",
		OID:     10104,
	},
	{
		Typname: "pg_sequence",
		OID:     10105,
	},
	{
		Typname: "_pg_publication",
		OID:     10106,
	},
	{
		Typname: "pg_publication",
		OID:     10107,
	},
	{
		Typname: "_pg_publication_namespace",
		OID:     10108,
	},
	{
		Typname: "pg_publication_namespace",
		OID:     10109,
	},
	{
		Typname: "_pg_publication_rel",
		OID:     10110,
	},
	{
		Typname: "pg_publication_rel",
		OID:     10111,
	},
	{
		Typname: "_pg_subscription",
		OID:     10112,
	},
	{
		Typname: "_pg_subscription_rel",
		OID:     10113,
	},
	{
		Typname: "pg_subscription_rel",
		OID:     10114,
	},
	{
		Typname: "_pg_roles",
		OID:     12001,
	},
	{
		Typname: "pg_roles",
		OID:     12002,
	},
	{
		Typname: "_pg_shadow",
		OID:     12006,
	},
	{
		Typname: "pg_shadow",
		OID:     12007,
	},
	{
		Typname: "_pg_group",
		OID:     12011,
	},
	{
		Typname: "pg_group",
		OID:     12012,
	},
	{
		Typname: "_pg_user",
		OID:     12015,
	},
	{
		Typname: "pg_user",
		OID:     12016,
	},
	{
		Typname: "_pg_policies",
		OID:     12019,
	},
	{
		Typname: "pg_policies",
		OID:     12020,
	},
	{
		Typname: "_pg_rules",
		OID:     12024,
	},
	{
		Typname: "pg_rules",
		OID:     12025,
	},
	{
		Typname: "_pg_views",
		OID:     12029,
	},
	{
		Typname: "pg_views",
		OID:     12030,
	},
	{
		Typname: "_pg_tables",
		OID:     12034,
	},
	{
		Typname: "pg_tables",
		OID:     12035,
	},
	{
		Typname: "_pg_matviews",
		OID:     12039,
	},
	{
		Typname: "pg_matviews",
		OID:     12040,
	},
	{
		Typname: "_pg_indexes",
		OID:     12044,
	},
	{
		Typname: "pg_indexes",
		OID:     12045,
	},
	{
		Typname: "_pg_sequences",
		OID:     12049,
	},
	{
		Typname: "pg_sequences",
		OID:     12050,
	},
	{
		Typname: "_pg_stats",
		OID:     12054,
	},
	{
		Typname: "pg_stats",
		OID:     12055,
	},
	{
		Typname: "_pg_stats_ext",
		OID:     12059,
	},
	{
		Typname: "pg_stats_ext",
		OID:     12060,
	},
	{
		Typname: "_pg_stats_ext_exprs",
		OID:     12064,
	},
	{
		Typname: "pg_stats_ext_exprs",
		OID:     12065,
	},
	{
		Typname: "_pg_publication_tables",
		OID:     12069,
	},
	{
		Typname: "pg_publication_tables",
		OID:     12070,
	},
	{
		Typname: "_pg_locks",
		OID:     12074,
	},
	{
		Typname: "pg_locks",
		OID:     12075,
	},
	{
		Typname: "_pg_cursors",
		OID:     12078,
	},
	{
		Typname: "pg_cursors",
		OID:     12079,
	},
	{
		Typname: "_pg_available_extensions",
		OID:     12082,
	},
	{
		Typname: "pg_available_extensions",
		OID:     12083,
	},
	{
		Typname: "_pg_available_extension_versions",
		OID:     12086,
	},
	{
		Typname: "pg_available_extension_versions",
		OID:     12087,
	},
	{
		Typname: "_pg_prepared_xacts",
		OID:     12091,
	},
	{
		Typname: "pg_prepared_xacts",
		OID:     12092,
	},
	{
		Typname: "_pg_prepared_statements",
		OID:     12096,
	},
	{
		Typname: "pg_prepared_statements",
		OID:     12097,
	},
	{
		Typname: "_pg_seclabels",
		OID:     12100,
	},
	{
		Typname: "pg_seclabels",
		OID:     12101,
	},
	{
		Typname: "_pg_settings",
		OID:     12105,
	},
	{
		Typname: "pg_settings",
		OID:     12106,
	},
	{
		Typname: "_pg_file_settings",
		OID:     12111,
	},
	{
		Typname: "pg_file_settings",
		OID:     12112,
	},
	{
		Typname: "_pg_hba_file_rules",
		OID:     12115,
	},
	{
		Typname: "pg_hba_file_rules",
		OID:     12116,
	},
	{
		Typname: "_pg_ident_file_mappings",
		OID:     12119,
	},
	{
		Typname: "pg_ident_file_mappings",
		OID:     12120,
	},
	{
		Typname: "_pg_timezone_abbrevs",
		OID:     12123,
	},
	{
		Typname: "pg_timezone_abbrevs",
		OID:     12124,
	},
	{
		Typname: "_pg_timezone_names",
		OID:     12127,
	},
	{
		Typname: "pg_timezone_names",
		OID:     12128,
	},
	{
		Typname: "_pg_config",
		OID:     12131,
	},
	{
		Typname: "pg_config",
		OID:     12132,
	},
	{
		Typname: "_pg_shmem_allocations",
		OID:     12135,
	},
	{
		Typname: "pg_shmem_allocations",
		OID:     12136,
	},
	{
		Typname: "_pg_backend_memory_contexts",
		OID:     12139,
	},
	{
		Typname: "pg_backend_memory_contexts",
		OID:     12140,
	},
	{
		Typname: "_pg_stat_all_tables",
		OID:     12143,
	},
	{
		Typname: "pg_stat_all_tables",
		OID:     12144,
	},
	{
		Typname: "_pg_stat_xact_all_tables",
		OID:     12148,
	},
	{
		Typname: "pg_stat_xact_all_tables",
		OID:     12149,
	},
	{
		Typname: "_pg_stat_sys_tables",
		OID:     12153,
	},
	{
		Typname: "pg_stat_sys_tables",
		OID:     12154,
	},
	{
		Typname: "_pg_stat_xact_sys_tables",
		OID:     12158,
	},
	{
		Typname: "pg_stat_xact_sys_tables",
		OID:     12159,
	},
	{
		Typname: "_pg_stat_user_tables",
		OID:     12162,
	},
	{
		Typname: "pg_stat_user_tables",
		OID:     12163,
	},
	{
		Typname: "_pg_stat_xact_user_tables",
		OID:     12167,
	},
	{
		Typname: "pg_stat_xact_user_tables",
		OID:     12168,
	},
	{
		Typname: "_pg_statio_all_tables",
		OID:     12171,
	},
	{
		Typname: "pg_statio_all_tables",
		OID:     12172,
	},
	{
		Typname: "_pg_statio_sys_tables",
		OID:     12176,
	},
	{
		Typname: "pg_statio_sys_tables",
		OID:     12177,
	},
	{
		Typname: "_pg_statio_user_tables",
		OID:     12180,
	},
	{
		Typname: "pg_statio_user_tables",
		OID:     12181,
	},
	{
		Typname: "_pg_stat_all_indexes",
		OID:     12184,
	},
	{
		Typname: "pg_stat_all_indexes",
		OID:     12185,
	},
	{
		Typname: "_pg_stat_sys_indexes",
		OID:     12189,
	},
	{
		Typname: "pg_stat_sys_indexes",
		OID:     12190,
	},
	{
		Typname: "_pg_stat_user_indexes",
		OID:     12193,
	},
	{
		Typname: "pg_stat_user_indexes",
		OID:     12194,
	},
	{
		Typname: "_pg_statio_all_indexes",
		OID:     12197,
	},
	{
		Typname: "pg_statio_all_indexes",
		OID:     12198,
	},
	{
		Typname: "_pg_statio_sys_indexes",
		OID:     12202,
	},
	{
		Typname: "pg_statio_sys_indexes",
		OID:     12203,
	},
	{
		Typname: "_pg_statio_user_indexes",
		OID:     12206,
	},
	{
		Typname: "pg_statio_user_indexes",
		OID:     12207,
	},
	{
		Typname: "_pg_statio_all_sequences",
		OID:     12210,
	},
	{
		Typname: "pg_statio_all_sequences",
		OID:     12211,
	},
	{
		Typname: "_pg_statio_sys_sequences",
		OID:     12215,
	},
	{
		Typname: "pg_statio_sys_sequences",
		OID:     12216,
	},
	{
		Typname: "_pg_statio_user_sequences",
		OID:     12219,
	},
	{
		Typname: "pg_statio_user_sequences",
		OID:     12220,
	},
	{
		Typname: "_pg_stat_activity",
		OID:     12223,
	},
	{
		Typname: "pg_stat_activity",
		OID:     12224,
	},
	{
		Typname: "_pg_stat_replication",
		OID:     12228,
	},
	{
		Typname: "pg_stat_replication",
		OID:     12229,
	},
	{
		Typname: "_pg_stat_slru",
		OID:     12233,
	},
	{
		Typname: "pg_stat_slru",
		OID:     12234,
	},
	{
		Typname: "_pg_stat_wal_receiver",
		OID:     12237,
	},
	{
		Typname: "pg_stat_wal_receiver",
		OID:     12238,
	},
	{
		Typname: "_pg_stat_recovery_prefetch",
		OID:     12241,
	},
	{
		Typname: "pg_stat_recovery_prefetch",
		OID:     12242,
	},
	{
		Typname: "_pg_stat_subscription",
		OID:     12245,
	},
	{
		Typname: "pg_stat_subscription",
		OID:     12246,
	},
	{
		Typname: "_pg_stat_ssl",
		OID:     12250,
	},
	{
		Typname: "pg_stat_ssl",
		OID:     12251,
	},
	{
		Typname: "_pg_stat_gssapi",
		OID:     12254,
	},
	{
		Typname: "pg_stat_gssapi",
		OID:     12255,
	},
	{
		Typname: "_pg_replication_slots",
		OID:     12258,
	},
	{
		Typname: "pg_replication_slots",
		OID:     12259,
	},
	{
		Typname: "_pg_stat_replication_slots",
		OID:     12263,
	},
	{
		Typname: "pg_stat_replication_slots",
		OID:     12264,
	},
	{
		Typname: "_pg_stat_database",
		OID:     12267,
	},
	{
		Typname: "pg_stat_database",
		OID:     12268,
	},
	{
		Typname: "_pg_stat_database_conflicts",
		OID:     12272,
	},
	{
		Typname: "pg_stat_database_conflicts",
		OID:     12273,
	},
	{
		Typname: "_pg_stat_user_functions",
		OID:     12276,
	},
	{
		Typname: "pg_stat_user_functions",
		OID:     12277,
	},
	{
		Typname: "_pg_stat_xact_user_functions",
		OID:     12281,
	},
	{
		Typname: "pg_stat_xact_user_functions",
		OID:     12282,
	},
	{
		Typname: "_pg_stat_archiver",
		OID:     12286,
	},
	{
		Typname: "pg_stat_archiver",
		OID:     12287,
	},
	{
		Typname: "_pg_stat_bgwriter",
		OID:     12290,
	},
	{
		Typname: "pg_stat_bgwriter",
		OID:     12291,
	},
	{
		Typname: "_pg_stat_wal",
		OID:     12294,
	},
	{
		Typname: "pg_stat_wal",
		OID:     12295,
	},
	{
		Typname: "_pg_stat_progress_analyze",
		OID:     12298,
	},
	{
		Typname: "pg_stat_progress_analyze",
		OID:     12299,
	},
	{
		Typname: "_pg_stat_progress_vacuum",
		OID:     12303,
	},
	{
		Typname: "pg_stat_progress_vacuum",
		OID:     12304,
	},
	{
		Typname: "_pg_stat_progress_cluster",
		OID:     12308,
	},
	{
		Typname: "pg_stat_progress_cluster",
		OID:     12309,
	},
	{
		Typname: "_pg_stat_progress_create_index",
		OID:     12313,
	},
	{
		Typname: "pg_stat_progress_create_index",
		OID:     12314,
	},
	{
		Typname: "_pg_stat_progress_basebackup",
		OID:     12318,
	},
	{
		Typname: "pg_stat_progress_basebackup",
		OID:     12319,
	},
	{
		Typname: "_pg_stat_progress_copy",
		OID:     12323,
	},
	{
		Typname: "pg_stat_progress_copy",
		OID:     12324,
	},
	{
		Typname: "_pg_user_mappings",
		OID:     12328,
	},
	{
		Typname: "pg_user_mappings",
		OID:     12329,
	},
	{
		Typname: "_pg_replication_origin_status",
		OID:     12333,
	},
	{
		Typname: "pg_replication_origin_status",
		OID:     12334,
	},
	{
		Typname: "_pg_stat_subscription_stats",
		OID:     12337,
	},
	{
		Typname: "pg_stat_subscription_stats",
		OID:     12338,
	},
	{
		Typname: "_cardinal_number",
		OID:     12416,
	},
	{
		Typname: "cardinal_number",
		OID:     12417,
	},
	{
		Typname: "_character_data",
		OID:     12419,
	},
	{
		Typname: "character_data",
		OID:     12420,
	},
	{
		Typname: "_sql_identifier",
		OID:     12421,
	},
	{
		Typname: "sql_identifier",
		OID:     12422,
	},
	{
		Typname: "_information_schema_catalog_name",
		OID:     12424,
	},
	{
		Typname: "information_schema_catalog_name",
		OID:     12425,
	},
	{
		Typname: "_time_stamp",
		OID:     12427,
	},
	{
		Typname: "time_stamp",
		OID:     12428,
	},
	{
		Typname: "_yes_or_no",
		OID:     12429,
	},
	{
		Typname: "yes_or_no",
		OID:     12430,
	},
	{
		Typname: "_applicable_roles",
		OID:     12433,
	},
	{
		Typname: "applicable_roles",
		OID:     12434,
	},
	{
		Typname: "_administrable_role_authorizations",
		OID:     12438,
	},
	{
		Typname: "administrable_role_authorizations",
		OID:     12439,
	},
	{
		Typname: "_attributes",
		OID:     12442,
	},
	{
		Typname: "attributes",
		OID:     12443,
	},
	{
		Typname: "_character_sets",
		OID:     12447,
	},
	{
		Typname: "character_sets",
		OID:     12448,
	},
	{
		Typname: "_check_constraint_routine_usage",
		OID:     12452,
	},
	{
		Typname: "check_constraint_routine_usage",
		OID:     12453,
	},
	{
		Typname: "_check_constraints",
		OID:     12457,
	},
	{
		Typname: "check_constraints",
		OID:     12458,
	},
	{
		Typname: "_collations",
		OID:     12462,
	},
	{
		Typname: "collations",
		OID:     12463,
	},
	{
		Typname: "_collation_character_set_applicability",
		OID:     12467,
	},
	{
		Typname: "collation_character_set_applicability",
		OID:     12468,
	},
	{
		Typname: "_column_column_usage",
		OID:     12472,
	},
	{
		Typname: "column_column_usage",
		OID:     12473,
	},
	{
		Typname: "_column_domain_usage",
		OID:     12477,
	},
	{
		Typname: "column_domain_usage",
		OID:     12478,
	},
	{
		Typname: "_column_privileges",
		OID:     12482,
	},
	{
		Typname: "column_privileges",
		OID:     12483,
	},
	{
		Typname: "_column_udt_usage",
		OID:     12487,
	},
	{
		Typname: "column_udt_usage",
		OID:     12488,
	},
	{
		Typname: "_columns",
		OID:     12492,
	},
	{
		Typname: "columns",
		OID:     12493,
	},
	{
		Typname: "_constraint_column_usage",
		OID:     12497,
	},
	{
		Typname: "constraint_column_usage",
		OID:     12498,
	},
	{
		Typname: "_constraint_table_usage",
		OID:     12502,
	},
	{
		Typname: "constraint_table_usage",
		OID:     12503,
	},
	{
		Typname: "_domain_constraints",
		OID:     12507,
	},
	{
		Typname: "domain_constraints",
		OID:     12508,
	},
	{
		Typname: "_domain_udt_usage",
		OID:     12512,
	},
	{
		Typname: "domain_udt_usage",
		OID:     12513,
	},
	{
		Typname: "_domains",
		OID:     12517,
	},
	{
		Typname: "domains",
		OID:     12518,
	},
	{
		Typname: "_enabled_roles",
		OID:     12522,
	},
	{
		Typname: "enabled_roles",
		OID:     12523,
	},
	{
		Typname: "_key_column_usage",
		OID:     12526,
	},
	{
		Typname: "key_column_usage",
		OID:     12527,
	},
	{
		Typname: "_parameters",
		OID:     12531,
	},
	{
		Typname: "parameters",
		OID:     12532,
	},
	{
		Typname: "_referential_constraints",
		OID:     12536,
	},
	{
		Typname: "referential_constraints",
		OID:     12537,
	},
	{
		Typname: "_role_column_grants",
		OID:     12541,
	},
	{
		Typname: "role_column_grants",
		OID:     12542,
	},
	{
		Typname: "_routine_column_usage",
		OID:     12545,
	},
	{
		Typname: "routine_column_usage",
		OID:     12546,
	},
	{
		Typname: "_routine_privileges",
		OID:     12550,
	},
	{
		Typname: "routine_privileges",
		OID:     12551,
	},
	{
		Typname: "_role_routine_grants",
		OID:     12555,
	},
	{
		Typname: "role_routine_grants",
		OID:     12556,
	},
	{
		Typname: "_routine_routine_usage",
		OID:     12559,
	},
	{
		Typname: "routine_routine_usage",
		OID:     12560,
	},
	{
		Typname: "_routine_sequence_usage",
		OID:     12564,
	},
	{
		Typname: "routine_sequence_usage",
		OID:     12565,
	},
	{
		Typname: "_routine_table_usage",
		OID:     12569,
	},
	{
		Typname: "routine_table_usage",
		OID:     12570,
	},
	{
		Typname: "_routines",
		OID:     12574,
	},
	{
		Typname: "routines",
		OID:     12575,
	},
	{
		Typname: "_schemata",
		OID:     12579,
	},
	{
		Typname: "schemata",
		OID:     12580,
	},
	{
		Typname: "_sequences",
		OID:     12583,
	},
	{
		Typname: "sequences",
		OID:     12584,
	},
	{
		Typname: "_sql_features",
		OID:     12588,
	},
	{
		Typname: "sql_features",
		OID:     12589,
	},
	{
		Typname: "_sql_implementation_info",
		OID:     12593,
	},
	{
		Typname: "sql_implementation_info",
		OID:     12594,
	},
	{
		Typname: "_sql_parts",
		OID:     12598,
	},
	{
		Typname: "sql_parts",
		OID:     12599,
	},
	{
		Typname: "_sql_sizing",
		OID:     12603,
	},
	{
		Typname: "sql_sizing",
		OID:     12604,
	},
	{
		Typname: "_table_constraints",
		OID:     12608,
	},
	{
		Typname: "table_constraints",
		OID:     12609,
	},
	{
		Typname: "_table_privileges",
		OID:     12613,
	},
	{
		Typname: "table_privileges",
		OID:     12614,
	},
	{
		Typname: "_role_table_grants",
		OID:     12618,
	},
	{
		Typname: "role_table_grants",
		OID:     12619,
	},
	{
		Typname: "_tables",
		OID:     12622,
	},
	{
		Typname: "tables",
		OID:     12623,
	},
	{
		Typname: "_transforms",
		OID:     12627,
	},
	{
		Typname: "transforms",
		OID:     12628,
	},
	{
		Typname: "_triggered_update_columns",
		OID:     12632,
	},
	{
		Typname: "triggered_update_columns",
		OID:     12633,
	},
	{
		Typname: "_triggers",
		OID:     12637,
	},
	{
		Typname: "triggers",
		OID:     12638,
	},
	{
		Typname: "_udt_privileges",
		OID:     12642,
	},
	{
		Typname: "udt_privileges",
		OID:     12643,
	},
	{
		Typname: "_role_udt_grants",
		OID:     12647,
	},
	{
		Typname: "role_udt_grants",
		OID:     12648,
	},
	{
		Typname: "_usage_privileges",
		OID:     12651,
	},
	{
		Typname: "usage_privileges",
		OID:     12652,
	},
	{
		Typname: "_role_usage_grants",
		OID:     12656,
	},
	{
		Typname: "role_usage_grants",
		OID:     12657,
	},
	{
		Typname: "_user_defined_types",
		OID:     12660,
	},
	{
		Typname: "user_defined_types",
		OID:     12661,
	},
	{
		Typname: "_view_column_usage",
		OID:     12665,
	},
	{
		Typname: "view_column_usage",
		OID:     12666,
	},
	{
		Typname: "_view_routine_usage",
		OID:     12670,
	},
	{
		Typname: "view_routine_usage",
		OID:     12671,
	},
	{
		Typname: "_view_table_usage",
		OID:     12675,
	},
	{
		Typname: "view_table_usage",
		OID:     12676,
	},
	{
		Typname: "_views",
		OID:     12680,
	},
	{
		Typname: "views",
		OID:     12681,
	},
	{
		Typname: "_data_type_privileges",
		OID:     12685,
	},
	{
		Typname: "data_type_privileges",
		OID:     12686,
	},
	{
		Typname: "_element_types",
		OID:     12690,
	},
	{
		Typname: "element_types",
		OID:     12691,
	},
	{
		Typname: "__pg_foreign_table_columns",
		OID:     12695,
	},
	{
		Typname: "_pg_foreign_table_columns",
		OID:     12696,
	},
	{
		Typname: "_column_options",
		OID:     12700,
	},
	{
		Typname: "column_options",
		OID:     12701,
	},
	{
		Typname: "__pg_foreign_data_wrappers",
		OID:     12704,
	},
	{
		Typname: "_pg_foreign_data_wrappers",
		OID:     12705,
	},
	{
		Typname: "_foreign_data_wrapper_options",
		OID:     12708,
	},
	{
		Typname: "foreign_data_wrapper_options",
		OID:     12709,
	},
	{
		Typname: "_foreign_data_wrappers",
		OID:     12712,
	},
	{
		Typname: "foreign_data_wrappers",
		OID:     12713,
	},
	{
		Typname: "__pg_foreign_servers",
		OID:     12716,
	},
	{
		Typname: "_pg_foreign_servers",
		OID:     12717,
	},
	{
		Typname: "_foreign_server_options",
		OID:     12721,
	},
	{
		Typname: "foreign_server_options",
		OID:     12722,
	},
	{
		Typname: "_foreign_servers",
		OID:     12725,
	},
	{
		Typname: "foreign_servers",
		OID:     12726,
	},
	{
		Typname: "__pg_foreign_tables",
		OID:     12729,
	},
	{
		Typname: "_pg_foreign_tables",
		OID:     12730,
	},
	{
		Typname: "_foreign_table_options",
		OID:     12734,
	},
	{
		Typname: "foreign_table_options",
		OID:     12735,
	},
	{
		Typname: "_foreign_tables",
		OID:     12738,
	},
	{
		Typname: "foreign_tables",
		OID:     12739,
	},
	{
		Typname: "__pg_user_mappings",
		OID:     12742,
	},
	{
		Typname: "_pg_user_mappings",
		OID:     12743,
	},
	{
		Typname: "_user_mapping_options",
		OID:     12747,
	},
	{
		Typname: "user_mapping_options",
		OID:     12748,
	},
	{
		Typname: "_user_mappings",
		OID:     12752,
	},
	{
		Typname: "user_mappings",
		OID:     12753,
	},
	{
		Typname: "_crypto_box_keypair",
		OID:     16657,
	},
	{
		Typname: "crypto_box_keypair",
		OID:     16658,
	},
	{
		Typname: "_crypto_sign_keypair",
		OID:     16664,
	},
	{
		Typname: "crypto_sign_keypair",
		OID:     16665,
	},
	{
		Typname: "_crypto_kx_keypair",
		OID:     16680,
	},
	{
		Typname: "crypto_kx_keypair",
		OID:     16681,
	},
	{
		Typname: "_crypto_kx_session",
		OID:     16686,
	},
	{
		Typname: "crypto_kx_session",
		OID:     16687,
	},
	{
		Typname: "_crypto_signcrypt_state_key",
		OID:     16754,
	},
	{
		Typname: "crypto_signcrypt_state_key",
		OID:     16755,
	},
	{
		Typname: "_crypto_signcrypt_keypair",
		OID:     16757,
	},
	{
		Typname: "crypto_signcrypt_keypair",
		OID:     16758,
	},
	{
		Typname: "_key_status",
		OID:     16771,
	},
	{
		Typname: "key_status",
		OID:     16772,
	},
	{
		Typname: "_key_type",
		OID:     16781,
	},
	{
		Typname: "key_type",
		OID:     16782,
	},
	{
		Typname: "_key",
		OID:     16789,
	},
	{
		Typname: "key",
		OID:     16790,
	},
	{
		Typname: "__key_id_context",
		OID:     16810,
	},
	{
		Typname: "_key_id_context",
		OID:     16811,
	},
	{
		Typname: "_valid_key",
		OID:     16893,
	},
	{
		Typname: "valid_key",
		OID:     16894,
	},
	{
		Typname: "_masking_rule",
		OID:     16909,
	},
	{
		Typname: "masking_rule",
		OID:     16910,
	},
	{
		Typname: "_mask_columns",
		OID:     16914,
	},
	{
		Typname: "mask_columns",
		OID:     16915,
	},
	{
		Typname: "_decrypted_key",
		OID:     16939,
	},
	{
		Typname: "decrypted_key",
		OID:     16940,
	},
	{
		Typname: "_pg_stat_statements_info",
		OID:     27035,
	},
	{
		Typname: "pg_stat_statements_info",
		OID:     27036,
	},
	{
		Typname: "_pg_stat_statements",
		OID:     27046,
	},
	{
		Typname: "pg_stat_statements",
		OID:     27047,
	},
	{
		Typname: "_aal_level",
		OID:     27105,
	},
	{
		Typname: "aal_level",
		OID:     27106,
	},
	{
		Typname: "_factor_status",
		OID:     27113,
	},
	{
		Typname: "factor_status",
		OID:     27114,
	},
	{
		Typname: "_factor_type",
		OID:     27119,
	},
	{
		Typname: "factor_type",
		OID:     27120,
	},
	{
		Typname: "_audit_log_entries",
		OID:     27142,
	},
	{
		Typname: "audit_log_entries",
		OID:     27143,
	},
	{
		Typname: "_identities",
		OID:     27148,
	},
	{
		Typname: "identities",
		OID:     27149,
	},
	{
		Typname: "_instances",
		OID:     27153,
	},
	{
		Typname: "instances",
		OID:     27154,
	},
	{
		Typname: "_mfa_amr_claims",
		OID:     27158,
	},
	{
		Typname: "mfa_amr_claims",
		OID:     27159,
	},
	{
		Typname: "_mfa_challenges",
		OID:     27163,
	},
	{
		Typname: "mfa_challenges",
		OID:     27164,
	},
	{
		Typname: "_mfa_factors",
		OID:     27168,
	},
	{
		Typname: "mfa_factors",
		OID:     27169,
	},
	{
		Typname: "_refresh_tokens",
		OID:     27173,
	},
	{
		Typname: "refresh_tokens",
		OID:     27174,
	},
	{
		Typname: "_saml_providers",
		OID:     27179,
	},
	{
		Typname: "saml_providers",
		OID:     27180,
	},
	{
		Typname: "_saml_relay_states",
		OID:     27187,
	},
	{
		Typname: "saml_relay_states",
		OID:     27188,
	},
	{
		Typname: "_schema_migrations",
		OID:     27193,
	},
	{
		Typname: "schema_migrations",
		OID:     27194,
	},
	{
		Typname: "_sessions",
		OID:     27196,
	},
	{
		Typname: "sessions",
		OID:     27197,
	},
	{
		Typname: "_sso_domains",
		OID:     27199,
	},
	{
		Typname: "sso_domains",
		OID:     27200,
	},
	{
		Typname: "_sso_providers",
		OID:     27205,
	},
	{
		Typname: "sso_providers",
		OID:     27206,
	},
	{
		Typname: "_users",
		OID:     27215,
	},
	{
		Typname: "users",
		OID:     27216,
	},
	{
		Typname: "_invoice",
		OID:     27228,
	},
	{
		Typname: "invoice",
		OID:     27229,
	},
	{
		Typname: "_subscription",
		OID:     27234,
	},
	{
		Typname: "subscription",
		OID:     27235,
	},
	{
		Typname: "_transactions",
		OID:     27240,
	},
	{
		Typname: "transactions",
		OID:     27241,
	},
	{
		Typname: "_buckets",
		OID:     27247,
	},
	{
		Typname: "buckets",
		OID:     27248,
	},
	{
		Typname: "_migrations",
		OID:     27255,
	},
	{
		Typname: "migrations",
		OID:     27256,
	},
	{
		Typname: "_objects",
		OID:     27259,
	},
	{
		Typname: "objects",
		OID:     27260,
	},
	{
		Typname: "_code_challenge_method",
		OID:     27516,
	},
	{
		Typname: "code_challenge_method",
		OID:     27517,
	},
	{
		Typname: "_flow_state",
		OID:     27522,
	},
	{
		Typname: "flow_state",
		OID:     27523,
	},
	{
		Typname: "_customers",
		OID:     27531,
	},
	{
		Typname: "customers",
		OID:     27532,
	},
}
