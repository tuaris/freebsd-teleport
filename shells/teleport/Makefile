PORTNAME=	teleport
DISTVERSIONPREFIX=	v
DISTVERSION=	17.3.0
CATEGORIES=	shells
PKGNAMEPREFIX=	gravitational-

MASTER_SITES+=	https://raw.githubusercontent.com/${GH_ACCOUNT}/${PORTNAME}/${DISTVERSIONFULL}/
DISTFILES=	go.mod 

MAINTAINER=	daniel@morante.net
COMMENT=	The easiest, and most secure way to access and protect all of your infrastructure. 
WWW=		https://github.com/gravitational/teleport

LICENSE=	AGPLv3
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:1.23,modules
USE_GITHUB=	yes
GH_ACCOUNT=	gravitational

FLAVORS=	client tctl
FLAVOR?=	${FLAVORS:[1]}

client_PKGNAMESUFFIX=	-client
client_PLIST_FILES=	bin/tsh
.if ${FLAVOR} == client
GO_TARGET=	./tool/tsh
.endif

tctl_PKGNAMESUFFIX=	-tctl
tctl_PLIST_FILES=	bin/tctl
.if ${FLAVOR} == tctl
GO_TARGET=	./tool/tctl
.endif

GO_MODULE=	github.com/gravitational/${PORTNAME}/v17

.include <bsd.port.mk>
