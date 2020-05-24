# Created by: Juraj Lutter <juraj@lutter.sk>
# $FreeBSD$

# XXX
# - backport fix from
#   https://people.freebsd.org/~dbaio/zabbix44-agent2-go-test.patch
#
# - cp zabbix50-server/files/zabbix_agentd.in zabbix50-server/files/zabbix_agent2d.in
#   adjust accordingly

PORTNAME=	zabbix50
CATEGORIES=	net-mgmt
PKGNAMESUFFIX=	-agent2
DISTFILES=	${DISTNAME}${EXTRACT_SUFX} # only to allow USE_GITHUB=nodefault

MASTERDIR=	${.CURDIR}/../zabbix50-server

EXTRA_PATCHES=	${.CURDIR}/files
DISTINFO_FILE=	${.CURDIR}/distinfo # only for GH_TUPLE

USES+=		autoreconf go:modules localbase pkgconfig iconv

# to generate gomod-vendor, uncomment the following line (or generate GH_TUPLE manually)
# WRKSRC=	${WRKDIR}/zabbix-${PORTVERSION}/src/go
USE_GITHUB=	nodefault
GH_ACCOUNT=	zabbix
GH_PROJECT=	zabbix
GH_TUPLE=	\
		alimy:mc:v2.0.3:alimy_mc_v2/src/go/vendor/github.com/alimy/mc/v2 \
		dustin:gomemcached:a2284a01c143:dustin_gomemcached/src/go/vendor/github.com/dustin/gomemcached \
		fsnotify:fsnotify:v1.4.7:fsnotify_fsnotify/src/go/vendor/github.com/fsnotify/fsnotify \
		go-ole:go-ole:v1.2.4:go_ole_go_ole/src/go/vendor/github.com/go-ole/go-ole \
		go-sql-driver:mysql:v1.5.0:go_sql_driver_mysql/src/go/vendor/github.com/go-sql-driver/mysql \
		godbus:dbus:v4.1.0:godbus_dbus/src/go/vendor/github.com/godbus/dbus \
		golang:sys:f43be2a4598c:golang_sys/src/go/vendor/golang.org/x/sys \
		golang:xerrors:1b5146add898:golang_xerrors/src/go/vendor/golang.org/x/xerrors \
		mattn:go-sqlite3:v2.0.3:mattn_go_sqlite3/src/go/vendor/github.com/mattn/go-sqlite3 \
		mediocregopher:radix:v3.4.2:mediocregopher_radix_v3/src/go/vendor/github.com/mediocregopher/radix/v3 \
		natefinch:npipe:c1b8fa8bdcce:natefinch_npipe/src/go/vendor/github.com/natefinch/npipe \
		shirou:gopsutil:v2.20.4:shirou_gopsutil/src/go/vendor/github.com/shirou/gopsutil

PLIST=	${PKGDIR}/pkg-plist.agent2

MAKE_ARGS+=	CGO_CFLAGS="-I${LOCALBASE}/include"

do-build:
	cd ${WRKSRC} && ${SETENV} ${MAKE_ENV} ${MAKE_CMD} ${MAKE_ARGS}

do-install:
	${MKDIR} ${STAGEDIR}${ETCDIR}
	${INSTALL_PROGRAM} ${WRKSRC}/src/go/bin/mock_server ${STAGEDIR}${LOCALBASE}/bin
	${INSTALL_PROGRAM} ${WRKSRC}/src/go/bin/zabbix_agent2 ${STAGEDIR}${LOCALBASE}/bin
	${CP} ${WRKSRC}/src/go/conf/zabbix_agent2.conf ${STAGEDIR}${ETCDIR}/zabbix_agent2.conf.sample
	${CP} ${WRKSRC}/src/go/conf/mock_server.conf ${STAGEDIR}${ETCDIR}/mock_server.conf.sample

.include "${MASTERDIR}/Makefile"
