# Created by: Juraj Lutter <juraj@lutter.sk>
# $FreeBSD$

PORTNAME=	zabbix50
CATEGORIES=	net-mgmt
PKGNAMESUFFIX=	-agent

MASTERDIR=	${.CURDIR}/../zabbix50-server
PLIST=		${PKGDIR}/pkg-plist.agent

OPTIONS_DEFINE=		CURL IPV6
OPTIONS_DEFAULT=	CURL OPENSSL

.include "${MASTERDIR}/Makefile"
