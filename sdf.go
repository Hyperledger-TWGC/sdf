/*
Copyright Hyperledger-TWGC All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
                 http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.


*/
package sdf

/*
#cgo windows CFLAGS: -DPACKED_STRUCTURES
#cgo linux LDFLAGS: -ldl
#cgo darwin LDFLAGS: -ldl
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <sansec/swsds.h>


SGD_HANDLE hDeviceHandle;
SGD_HANDLE hSessionHandle;

typedef unsigned char     SGD_UCHAR;
typedef unsigned char*    SGD_UCHAR_PRT;


#ifdef _WIN32
#include<windows.h>

struct LibHandle {
	HMODULE handle;
};



struct LibHandle *New(const char *iLibrary)
{
	struct LibHandle *h = calloc(1,sizeof(struct LibHandle));
	h->handle = LoadLibrary(iLibrary);
	if (h->handle == NULL) {
		free(h);
		return NULL;
	}

	return h;
}

void Destroy(struct LibHandle *h)
{
	if(!h){
		return ;
	}
    if (h->handle == NULL) {
		return;
	}
	free(h);

}

#else
#include <dlfcn.h>

struct LibHandle {
	void *handle;
};

struct LibHandle *New(const char *iLibrary)
{

	struct LibHandle *h = calloc(1,sizeof(struct LibHandle));
	h->handle = dlopen(iLibrary,1);
	if(h->handle == NULL){
		free(h);
		return NULL;
	}
	return h;
}




void Destroy(struct LibHandle *h)
{
	if (!h) {
		return;
	}
	if (h->handle == NULL) {
		return;
	}
	if (dlclose(h->handle) < 0) {
		return;
	}
	free(h);
}

#endif

SGD_RV SDFOpenDevice(struct LibHandle * h,SGD_HANDLE *phDeviceHandle)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE*);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_OpenDevice");
	return (*fptr)(phDeviceHandle);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_OpenDevice");
	return (*fptr)(phDeviceHandle);
#endif
}

SGD_RV SDFCloseDevice(struct LibHandle * h,SGD_HANDLE hDeviceHandle)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_CloseDevice");
	return (*fptr)(hDeviceHandle);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_CloseDevice");
	return (*fptr)(hDeviceHandle);
#endif

}

SGD_RV SDFOpenSession(struct LibHandle * h,SGD_HANDLE hDeviceHandle, SGD_HANDLE *phSessionHandle)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE,SGD_HANDLE *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_OpenSession");
	return (*fptr)(hDeviceHandle,phSessionHandle);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_OpenSession");
	return (*fptr)(hDeviceHandle,phSessionHandle);
#endif

}

SGD_RV SDFCloseSession(struct LibHandle * h,SGD_HANDLE hSessionHandle)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_CloseSession");
	return (*fptr)(hSessionHandle);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_CloseSession");
	return (*fptr)(hSessionHandle);
#endif

}

SGD_RV SDFGetDeviceInfo(struct LibHandle * h,SGD_HANDLE hSessionHandle, DEVICEINFO *pstDeviceInfo)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE,DEVICEINFO *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_GetDeviceInfo");
	return (*fptr)(hSessionHandle,pstDeviceInfo);
#else

	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_GetDeviceInfo");
	return (*fptr)(hSessionHandle,pstDeviceInfo);
#endif
}

SGD_RV SDFGenerateRandom(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32  uiLength, SGD_UCHAR_PRT *pucRandom)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE,SGD_UINT32,SGD_UCHAR*);
	*pucRandom = calloc(uiLength, sizeof(SGD_UCHAR));
	if (*pucRandom == NULL) {
		return SGD_FALSE;
	}
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_GenerateRandom");
	return (*fptr)(hSessionHandle,uiLength,*pucRandom);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_GenerateRandom");
	return (*fptr)(hSessionHandle,uiLength,*pucRandom);
#endif
}

SGD_RV SDFGetPrivateKeyAccessRight(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32 uiKeyIndex,SGD_UCHAR_PRT pucPassword, SGD_UINT32  uiPwdLength)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE,SGD_UINT32,SGD_UCHAR*,SGD_UINT32);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_GetPrivateKeyAccessRight");
	return (*fptr)(hSessionHandle,uiKeyIndex,pucPassword,uiPwdLength);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_GetPrivateKeyAccessRight");
	return (*fptr)(hSessionHandle,uiKeyIndex,pucPassword,uiPwdLength);
#endif
}

SGD_RV SDFReleasePrivateKeyAccessRight(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE,SGD_UINT32);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ReleasePrivateKeyAccessRight");
	return (*fptr)(hSessionHandle,uiKeyIndex);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ReleasePrivateKeyAccessRight");
	return (*fptr)(hSessionHandle,uiKeyIndex);
#endif
}

SGD_RV SDFExportSignPublicKey_RSA(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex,RSArefPublicKey *pucPublicKey)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE,SGD_UINT32,RSArefPublicKey*);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ExportSignPublicKey_RSA");
	return (*fptr)(hSessionHandle,uiKeyIndex,pucPublicKey);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ExportSignPublicKey_RSA");
	return (*fptr)(hSessionHandle,uiKeyIndex,pucPublicKey);
#endif
}
SGD_RV SDFExportEncPublicKey_RSA(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex,RSArefPublicKey *pucPublicKey)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE,SGD_UINT32,RSArefPublicKey*);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ExportEncPublicKey_RSA");
	return (*fptr)(hSessionHandle,uiKeyIndex,pucPublicKey);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ExportEncPublicKey_RSA");
	return (*fptr)(hSessionHandle,uiKeyIndex,pucPublicKey);
#endif
}
SGD_RV SDFGenerateKeyPair_RSA(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyBits,RSArefPublicKey *pucPublicKey,RSArefPrivateKey *pucPrivateKey)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE,SGD_UINT32,RSArefPublicKey*,RSArefPrivateKey*);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_GenerateKeyPair_RSA");
	return (*fptr)(hSessionHandle,uiKeyBits,pucPublicKey,pucPrivateKey);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_GenerateKeyPair_RSA");
	return (*fptr)(hSessionHandle,uiKeyBits,pucPublicKey,pucPrivateKey);
#endif
}
SGD_RV SDFGenerateKeyWithIPK_RSA(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32 uiIPKIndex,SGD_UINT32 uiKeyBits,SGD_UCHAR_PRT *pucKey,SGD_UINT32 *puiKeyLength,SGD_HANDLE *phKeyHandle)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE,SGD_UINT32,SGD_UINT32,SGD_UCHAR*,SGD_UINT32*,SGD_HANDLE*);
	*pucKey = calloc(*puiKeyLength, sizeof(SGD_UCHAR));
	if (*pucKey == NULL) {
		return SGD_FALSE;
	}
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_GenerateKeyWithIPK_RSA");
	return (*fptr)(hSessionHandle,uiIPKIndex,uiKeyBits,*pucKey,puiKeyLength,phKeyHandle);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_GenerateKeyWithIPK_RSA");
	return (*fptr)(hSessionHandle,uiIPKIndex,uiKeyBits,*pucKey,puiKeyLength,phKeyHandle);
#endif
}
SGD_RV SDFGenerateKeyWithEPK_RSA(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32 uiKeyBits,RSArefPublicKey *pucPublicKey,SGD_UCHAR_PRT *pucKey,SGD_UINT32 *puiKeyLength,SGD_HANDLE *phKeyHandle)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE,SGD_UINT32,RSArefPublicKey*,SGD_UCHAR*,SGD_UINT32*,SGD_HANDLE*);
	*pucKey = calloc(*puiKeyLength, sizeof(SGD_UCHAR));
	if (*pucKey == NULL) {
		return SGD_FALSE;
	}
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_GenerateKeyWithEPK_RSA");
	return (*fptr)(hSessionHandle,uiKeyBits,pucPublicKey,*pucKey,puiKeyLength,phKeyHandle);
#else
	FPTR fptr = (FPTR)dlsym(h->handle,"SDF_GenerateKeyWithEPK_RSA");
	return (*fptr)(hSessionHandle,uiKeyBits,pucPublicKey,*pucKey,puiKeyLength,phKeyHandle);
#endif
}

SGD_RV SDFImportKeyWithISK_RSA(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32 uiISKIndex,SGD_UCHAR_PRT *pucKey,SGD_UINT32 uiKeyLength,SGD_HANDLE *phKeyHandle)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE,SGD_UINT32,SGD_UCHAR*,SGD_UINT32,SGD_HANDLE*);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ImportKeyWithISK_RSA");
	return (*fptr)(hSessionHandle,uiISKIndex,*pucKey,uiKeyLength,phKeyHandle);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ImportKeyWithISK_RSA");
	return (*fptr)(hSessionHandle,uiISKIndex,*pucKey,uiKeyLength,phKeyHandle);
#endif
}
SGD_RV SDFExchangeDigitEnvelopeBaseOnRSA(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex,RSArefPublicKey *pucPublicKey,SGD_UCHAR_PRT *pucDEInput,SGD_UINT32  uiDELength,SGD_UCHAR_PRT *pucDEOutput,SGD_UINT32  *puiDELength)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE,SGD_UINT32,RSArefPublicKey*,SGD_UCHAR*,SGD_UINT32,SGD_UCHAR*,SGD_UINT32*);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ExchangeDigitEnvelopeBaseOnRSA");
	return (*fptr)(hSessionHandle,uiKeyIndex,pucPublicKey,*pucDEInput,uiDELength,*pucDEOutput,puiDELength);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ExchangeDigitEnvelopeBaseOnRSA");
	return (*fptr)(hSessionHandle,uiKeyIndex,pucPublicKey,*pucDEInput,uiDELength,*pucDEOutput,puiDELength);
#endif
}
SGD_RV SDFExportSignPublicKey_ECC(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex,ECCrefPublicKey *pucPublicKey)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE,SGD_UINT32,ECCrefPublicKey*);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ExportSignPublicKey_ECC");
	return (*fptr)(hSessionHandle,uiKeyIndex,pucPublicKey);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ExportSignPublicKey_ECC");
	return (*fptr)(hSessionHandle,uiKeyIndex,pucPublicKey);
#endif
}
SGD_RV SDFExportEncPublicKey_ECC(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex,ECCrefPublicKey *pucPublicKey)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE,SGD_UINT32,ECCrefPublicKey*);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ExportEncPublicKey_ECC");
	return (*fptr)(hSessionHandle,uiKeyIndex,pucPublicKey);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ExportEncPublicKey_ECC");
	return (*fptr)(hSessionHandle,uiKeyIndex,pucPublicKey);
#endif
}
SGD_RV SDFGenerateKeyPair_ECC(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32  uiAlgID,SGD_UINT32  uiKeyBits,ECCrefPublicKey *pucPublicKey,ECCrefPrivateKey *pucPrivateKey)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE,SGD_UINT32,SGD_UINT32,ECCrefPublicKey*,ECCrefPrivateKey*);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_GenerateKeyPair_ECC");
	return (*fptr)(hSessionHandle,uiAlgID,uiKeyBits,pucPublicKey,pucPrivateKey);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_GenerateKeyPair_ECC");
	return (*fptr)(hSessionHandle,uiAlgID,uiKeyBits,pucPublicKey,pucPrivateKey);
#endif
}
SGD_RV SDFGenerateKeyWithIPK_ECC (struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32 uiIPKIndex,SGD_UINT32 uiKeyBits,ECCCipher *pucKey,SGD_HANDLE *phKeyHandle)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE,SGD_UINT32,SGD_UINT32,ECCCipher*,SGD_HANDLE*);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_GenerateKeyWithIPK_ECC");
	return (*fptr)(hSessionHandle,uiIPKIndex,uiKeyBits,pucKey,phKeyHandle);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_GenerateKeyWithIPK_ECC");
	return (*fptr)(hSessionHandle,uiIPKIndex,uiKeyBits,pucKey,phKeyHandle);
#endif
}
SGD_RV SDFGenerateKeyWithEPK_ECC (struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32 uiKeyBits,SGD_UINT32  uiAlgID,ECCrefPublicKey *pucPublicKey,ECCCipher *pucKey,SGD_HANDLE *phKeyHandle)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE,SGD_UINT32,SGD_UINT32,ECCrefPublicKey*,ECCCipher*,SGD_HANDLE*);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_GenerateKeyWithEPK_ECC");
	return (*fptr)(hSessionHandle,uiKeyBits,uiAlgID,pucPublicKey,pucKey,phKeyHandle);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_GenerateKeyWithEPK_ECC");
	return (*fptr)(hSessionHandle,uiKeyBits,uiAlgID,pucPublicKey,pucKey,phKeyHandle);
#endif
}
SGD_RV SDFImportKeyWithISK_ECC (struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UINT32 uiISKIndex,ECCCipher *pucKey,SGD_HANDLE *phKeyHandle)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE,SGD_UINT32,ECCCipher*,SGD_HANDLE*);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ImportKeyWithISK_ECC");
	return (*fptr)(hSessionHandle,uiISKIndex,pucKey,phKeyHandle);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ImportKeyWithISK_ECC");
	return (*fptr)(hSessionHandle,uiISKIndex,pucKey,phKeyHandle);
#endif
}
SGD_RV SDFGenerateAgreementDataWithECC (struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32 uiISKIndex,SGD_UINT32 uiKeyBits,SGD_UCHAR_PRT *pucSponsorID,SGD_UINT32 uiSponsorIDLength,ECCrefPublicKey  *pucSponsorPublicKey,ECCrefPublicKey  *pucSponsorTmpPublicKey,SGD_HANDLE *phAgreementHandle)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE , SGD_UINT32 ,SGD_UINT32 ,SGD_UCHAR *,SGD_UINT32 ,ECCrefPublicKey  *,ECCrefPublicKey  *,SGD_HANDLE *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_GenerateAgreementDataWithECC");
	return (*fptr)(hSessionHandle,  uiISKIndex, uiKeyBits, *pucSponsorID, uiSponsorIDLength,  pucSponsorPublicKey,  pucSponsorTmpPublicKey, phAgreementHandle);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_GenerateAgreementDataWithECC");
	return (*fptr)(hSessionHandle,  uiISKIndex, uiKeyBits, *pucSponsorID, uiSponsorIDLength,  pucSponsorPublicKey,  pucSponsorTmpPublicKey, phAgreementHandle);
#endif
}
SGD_RV SDFGenerateKeyWithECC (struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UCHAR_PRT *pucResponseID,SGD_UINT32 uiResponseIDLength,ECCrefPublicKey *pucResponsePublicKey,ECCrefPublicKey *pucResponseTmpPublicKey,SGD_HANDLE hAgreementHandle,SGD_HANDLE *phKeyHandle)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE, SGD_UCHAR *,SGD_UINT32 ,ECCrefPublicKey *,ECCrefPublicKey *,SGD_HANDLE ,SGD_HANDLE *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_GenerateKeyWithECC");
	return (*fptr)(hSessionHandle,*pucResponseID,uiResponseIDLength,pucResponsePublicKey,pucResponseTmpPublicKey,hAgreementHandle,phKeyHandle);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_GenerateKeyWithECC");
	return (*fptr)(hSessionHandle,*pucResponseID,uiResponseIDLength,pucResponsePublicKey,pucResponseTmpPublicKey,hAgreementHandle,phKeyHandle);
#endif
}
SGD_RV SDFGenerateAgreementDataAndKeyWithECC (struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32 uiISKIndex,SGD_UINT32 uiKeyBits,SGD_UCHAR_PRT *pucResponseID,SGD_UINT32 uiResponseIDLength,SGD_UCHAR_PRT *pucSponsorID,SGD_UINT32 uiSponsorIDLength,ECCrefPublicKey *pucSponsorPublicKey,ECCrefPublicKey *pucSponsorTmpPublicKey,ECCrefPublicKey  *pucResponsePublicKey,	ECCrefPublicKey  *pucResponseTmpPublicKey,SGD_HANDLE *phKeyHandle)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE , SGD_UINT32 ,SGD_UINT32 ,SGD_UCHAR *,SGD_UINT32 ,SGD_UCHAR *,SGD_UINT32 ,ECCrefPublicKey *,ECCrefPublicKey *,ECCrefPublicKey  *,	ECCrefPublicKey  *,SGD_HANDLE *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_GenerateAgreementDataAndKeyWithECC");
	return (*fptr)(hSessionHandle,  uiISKIndex, uiKeyBits, *pucResponseID, uiResponseIDLength, *pucSponsorID, uiSponsorIDLength, pucSponsorPublicKey, pucSponsorTmpPublicKey,  pucResponsePublicKey,	  pucResponseTmpPublicKey, phKeyHandle);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_GenerateAgreementDataAndKeyWithECC");
	return (*fptr)(hSessionHandle,  uiISKIndex, uiKeyBits, *pucResponseID, uiResponseIDLength, *pucSponsorID, uiSponsorIDLength, pucSponsorPublicKey, pucSponsorTmpPublicKey,  pucResponsePublicKey,	  pucResponseTmpPublicKey, phKeyHandle);
#endif
}
SGD_RV SDFExchangeDigitEnvelopeBaseOnECC(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32  uiKeyIndex,SGD_UINT32  uiAlgID,ECCrefPublicKey *pucPublicKey,ECCCipher *pucEncDataIn,ECCCipher *pucEncDataOut)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE , SGD_UINT32  ,SGD_UINT32  ,ECCrefPublicKey *,ECCCipher *,ECCCipher *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ExchangeDigitEnvelopeBaseOnECC");
	return (*fptr)(hSessionHandle,   uiKeyIndex,  uiAlgID, pucPublicKey, pucEncDataIn, pucEncDataOut);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ExchangeDigitEnvelopeBaseOnECC");
	return (*fptr)(hSessionHandle,   uiKeyIndex,  uiAlgID, pucPublicKey, pucEncDataIn, pucEncDataOut);
#endif
}
SGD_RV SDFGenerateKeyWithKEK(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32 uiKeyBits,SGD_UINT32  uiAlgID,SGD_UINT32 uiKEKIndex, SGD_UCHAR_PRT *pucKey, SGD_UINT32 *puiKeyLength, SGD_HANDLE *phKeyHandle)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE , SGD_UINT32 ,SGD_UINT32  ,SGD_UINT32 , SGD_UCHAR *, SGD_UINT32 *, SGD_HANDLE *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_GenerateKeyWithKEK");
	return (*fptr)(hSessionHandle,  uiKeyBits,  uiAlgID, uiKEKIndex,  *pucKey,  puiKeyLength,  phKeyHandle);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_GenerateKeyWithKEK");
	return (*fptr)(hSessionHandle,  uiKeyBits,  uiAlgID, uiKEKIndex,  *pucKey,  puiKeyLength,  phKeyHandle);
#endif
}
SGD_RV SDFImportKeyWithKEK(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32  uiAlgID,SGD_UINT32 uiKEKIndex, SGD_UCHAR_PRT *pucKey, SGD_UINT32 uiKeyLength, SGD_HANDLE *phKeyHandle)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE , SGD_UINT32  ,SGD_UINT32 , SGD_UCHAR *, SGD_UINT32 , SGD_HANDLE *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ImportKeyWithKEK");
	return (*fptr)(hSessionHandle,  uiAlgID, uiKEKIndex,  *pucKey,  uiKeyLength,  phKeyHandle);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ImportKeyWithKEK");
	return (*fptr)(hSessionHandle,  uiAlgID, uiKEKIndex,  *pucKey,  uiKeyLength,  phKeyHandle);
#endif
}
SGD_RV SDFDestroyKey(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_HANDLE hKeyHandle)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE , SGD_HANDLE);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_DestroyKey");
	return (*fptr)(hSessionHandle,  hKeyHandle);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_DestroyKey");
	return (*fptr)(hSessionHandle,  hKeyHandle);
#endif
}



SGD_RV SDFExternalPublicKeyOperation_RSA(struct LibHandle * h,SGD_HANDLE hSessionHandle, RSArefPublicKey *pucPublicKey,SGD_UCHAR_PRT *pucDataInput,SGD_UINT32  uiInputLength,SGD_UCHAR_PRT *pucDataOutput,SGD_UINT32  *puiOutputLength)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE , RSArefPublicKey *,SGD_UCHAR *,SGD_UINT32  ,SGD_UCHAR *,SGD_UINT32  *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ExternalPublicKeyOperation_RSA");
	return (*fptr)(hSessionHandle,  pucPublicKey, *pucDataInput,  uiInputLength, *pucDataOutput,  puiOutputLength);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ExternalPublicKeyOperation_RSA");
	return (*fptr)(hSessionHandle,  pucPublicKey, *pucDataInput,  uiInputLength, *pucDataOutput,  puiOutputLength);
#endif
}

SGD_RV SDFInternalPublicKeyOperation_RSA(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UINT32  uiKeyIndex,SGD_UCHAR_PRT pucDataInput,SGD_UINT32  uiInputLength,SGD_UCHAR_PRT *pucDataOutput,SGD_UINT32  *puiOutputLength)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UINT32  ,SGD_UCHAR *,SGD_UINT32  ,SGD_UCHAR *,SGD_UINT32  *);
    *pucDataOutput = calloc(*puiOutputLength, sizeof(SGD_UCHAR));
	if (*pucDataOutput == NULL) {
		return SGD_FALSE;
	}
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_InternalPublicKeyOperation_RSA");
	return (*fptr)(hSessionHandle,  uiKeyIndex, pucDataInput,  uiInputLength, *pucDataOutput,  puiOutputLength);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_InternalPublicKeyOperation_RSA");
	return (*fptr)(hSessionHandle,  uiKeyIndex, pucDataInput,  uiInputLength, *pucDataOutput,  puiOutputLength);
#endif
}
SGD_RV SDFInternalPrivateKeyOperation_RSA(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UINT32  uiKeyIndex,SGD_UCHAR_PRT pucDataInput,SGD_UINT32  uiInputLength,SGD_UCHAR_PRT *pucDataOutput,SGD_UINT32  *puiOutputLength)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UINT32  ,SGD_UCHAR *,SGD_UINT32  ,SGD_UCHAR *,SGD_UINT32  *);
    *pucDataOutput = calloc(*puiOutputLength, sizeof(SGD_UCHAR));
	if (*pucDataOutput == NULL) {
		return SGD_FALSE;
	}
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_InternalPrivateKeyOperation_RSA");
	return (*fptr)(hSessionHandle,  uiKeyIndex, pucDataInput,uiInputLength,*pucDataOutput,puiOutputLength);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_InternalPrivateKeyOperation_RSA");
	return (*fptr)(hSessionHandle,  uiKeyIndex, pucDataInput,uiInputLength,*pucDataOutput,puiOutputLength);
#endif
}

SGD_RV SDFExternalVerify_ECC(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UINT32 uiAlgID,ECCrefPublicKey *pucPublicKey,SGD_UCHAR_PRT *pucDataInput,SGD_UINT32  uiInputLength,ECCSignature *pucSignature)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UINT32 ,ECCrefPublicKey *,SGD_UCHAR *,SGD_UINT32  ,ECCSignature *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ExternalVerify_ECC");
	return (*fptr)(hSessionHandle, uiAlgID, pucPublicKey, *pucDataInput,  uiInputLength, pucSignature);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ExternalVerify_ECC");
	return (*fptr)(hSessionHandle, uiAlgID, pucPublicKey, *pucDataInput,  uiInputLength, pucSignature);
#endif
}
SGD_RV SDFInternalSign_ECC(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UINT32  uiISKIndex,SGD_UCHAR_PRT pucData,SGD_UINT32  uiDataLength,ECCSignature *pucSignature)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UINT32  uiISKIndex,SGD_UCHAR *pucData,SGD_UINT32  uiDataLength,ECCSignature *pucSignature);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_InternalSign_ECC");
	return (*fptr)(hSessionHandle,  uiISKIndex, pucData,  uiDataLength, pucSignature);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_InternalSign_ECC");
	return (*fptr)(hSessionHandle,  uiISKIndex, pucData,  uiDataLength, pucSignature);
#endif
}
SGD_RV SDFInternalVerify_ECC(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UINT32  uiISKIndex,SGD_UCHAR_PRT pucData,SGD_UINT32  uiDataLength,ECCSignature *pucSignature)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UINT32  ,SGD_UCHAR *,SGD_UINT32  ,ECCSignature *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_InternalVerify_ECC");
	return (*fptr)(hSessionHandle,  uiISKIndex, pucData,  uiDataLength, pucSignature);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_InternalVerify_ECC");
	return (*fptr)(hSessionHandle,  uiISKIndex, pucData,  uiDataLength, pucSignature);
#endif
}
SGD_RV SDFExternalEncrypt_ECC(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UINT32 uiAlgID,ECCrefPublicKey *pucPublicKey,SGD_UCHAR_PRT pucData,SGD_UINT32  uiDataLength,ECCCipher *pucEncData)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UINT32 ,ECCrefPublicKey *,SGD_UCHAR *,SGD_UINT32  ,ECCCipher *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ExternalEncrypt_ECC");
	return (*fptr)(hSessionHandle, uiAlgID, pucPublicKey, pucData,  uiDataLength, pucEncData);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ExternalEncrypt_ECC");
	return (*fptr)(hSessionHandle, uiAlgID, pucPublicKey, pucData,  uiDataLength, pucEncData);
#endif
}

SGD_RV SDFEncrypt(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_HANDLE hKeyHandle,SGD_UINT32 uiAlgID,SGD_UCHAR_PRT pucIV,SGD_UCHAR_PRT pucData,SGD_UINT32 uiDataLength,SGD_UCHAR_PRT *pucEncData,SGD_UINT32  *puiEncDataLength)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_HANDLE ,SGD_UINT32 ,SGD_UCHAR *,SGD_UCHAR *,SGD_UINT32 ,SGD_UCHAR *,SGD_UINT32  *);
    *pucEncData = calloc(*puiEncDataLength, sizeof(SGD_UCHAR));
	if (*pucEncData == NULL) {
		return SGD_FALSE;
	}
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_Encrypt");
	return (*fptr)(hSessionHandle, hKeyHandle, uiAlgID, pucIV, pucData, uiDataLength, *pucEncData,  puiEncDataLength);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_Encrypt");
	return (*fptr)(hSessionHandle, hKeyHandle, uiAlgID, pucIV, pucData, uiDataLength, *pucEncData,  puiEncDataLength);
#endif
}
SGD_RV SDFDecrypt (struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_HANDLE hKeyHandle,SGD_UINT32 uiAlgID,SGD_UCHAR_PRT pucIV,SGD_UCHAR_PRT pucEncData,SGD_UINT32  uiEncDataLength,SGD_UCHAR_PRT *pucData,SGD_UINT32 *puiDataLength)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_HANDLE ,SGD_UINT32 ,SGD_UCHAR *,SGD_UCHAR *,SGD_UINT32  ,SGD_UCHAR *,SGD_UINT32 *);
    *pucData = calloc(*puiDataLength, sizeof(SGD_UCHAR));
	if (*pucData == NULL) {
		return SGD_FALSE;
	}
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_Decrypt");
	return (*fptr)(hSessionHandle, hKeyHandle, uiAlgID, pucIV, pucEncData,  uiEncDataLength, *pucData, puiDataLength);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_Decrypt");
	return (*fptr)(hSessionHandle, hKeyHandle, uiAlgID, pucIV, pucEncData,  uiEncDataLength, *pucData, puiDataLength);
#endif
}
SGD_RV SDFCalculateMAC(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_HANDLE hKeyHandle,SGD_UINT32 uiAlgID,SGD_UCHAR_PRT pucIV,SGD_UCHAR_PRT pucData,SGD_UINT32 uiDataLength,SGD_UCHAR_PRT *pucMAC,SGD_UINT32  *puiMACLength)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_HANDLE ,SGD_UINT32 ,SGD_UCHAR *,SGD_UCHAR *,SGD_UINT32 ,SGD_UCHAR *,SGD_UINT32  *);
	*pucMAC = calloc(*puiMACLength, sizeof(SGD_UCHAR));
	if (*pucMAC == NULL) {
		return SGD_FALSE;
	}
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_CalculateMAC");
	return (*fptr)(hSessionHandle, hKeyHandle, uiAlgID, pucIV, pucData, uiDataLength, *pucMAC,  puiMACLength);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_CalculateMAC");
	return (*fptr)(hSessionHandle, hKeyHandle, uiAlgID, pucIV, pucData, uiDataLength, *pucMAC,  puiMACLength);
#endif
}


SGD_RV SDFCreateFile(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UCHAR_PRT pucFileName,SGD_UINT32 uiNameLen,SGD_UINT32 uiFileSize)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE,SGD_UCHAR *,SGD_UINT32 ,SGD_UINT32 );
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_CreateFile");
	return (*fptr)(hSessionHandle, pucFileName, uiNameLen, uiFileSize);
#else

	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_CreateFile");
	return (*fptr)(hSessionHandle, pucFileName, uiNameLen, uiFileSize);

#endif
}
SGD_RV SDFReadFile(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UCHAR_PRT pucFileName,SGD_UINT32 uiNameLen,SGD_UINT32 uiOffset,SGD_UINT32 *puiReadLength,SGD_UCHAR_PRT *pucBuffer)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UCHAR *,SGD_UINT32 ,SGD_UINT32 ,SGD_UINT32 *,SGD_UCHAR *);
	*pucBuffer = calloc(*puiReadLength, sizeof(SGD_UCHAR));
	if (*pucBuffer == NULL) {
		return SGD_FALSE;
	}
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ReadFile");
	return (*fptr)(hSessionHandle, pucFileName, uiNameLen, uiOffset, puiReadLength, *pucBuffer);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ReadFile");
	return (*fptr)(hSessionHandle, pucFileName, uiNameLen, uiOffset, puiReadLength, *pucBuffer);
#endif
}
SGD_RV SDFWriteFile(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UCHAR_PRT pucFileName,SGD_UINT32 uiNameLen,SGD_UINT32 uiOffset,SGD_UINT32 uiWriteLength,SGD_UCHAR_PRT pucBuffer)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UCHAR *,SGD_UINT32 ,SGD_UINT32 ,SGD_UINT32 ,SGD_UCHAR *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_WriteFile");
	return (*fptr)(hSessionHandle, pucFileName, uiNameLen, uiOffset, uiWriteLength, pucBuffer);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_WriteFile");
	return (*fptr)(hSessionHandle, pucFileName, uiNameLen, uiOffset, uiWriteLength, pucBuffer);
#endif
}
SGD_RV SDFDeleteFile(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UCHAR_PRT pucFileName,SGD_UINT32 uiNameLen)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UCHAR *,SGD_UINT32 );
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_DeleteFile");
	return (*fptr)(hSessionHandle, pucFileName, uiNameLen);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_DeleteFile");
	return (*fptr)(hSessionHandle, pucFileName, uiNameLen);
#endif
}


SGD_RV SDFHashInit(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UINT32 uiAlgID,ECCrefPublicKey *pucPublicKey,SGD_UCHAR_PRT pucID,SGD_UINT32 uiIDLength)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UINT32 ,ECCrefPublicKey *,SGD_UCHAR *,SGD_UINT32 );
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_HashInit");
	return (*fptr)(hSessionHandle, uiAlgID, pucPublicKey, pucID, uiIDLength);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_HashInit");
	return (*fptr)(hSessionHandle, uiAlgID, pucPublicKey, pucID, uiIDLength);
#endif
}
SGD_RV SDFHashUpdate(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UCHAR_PRT pucData,SGD_UINT32  uiDataLength)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UCHAR *,SGD_UINT32  );
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_HashUpdate");
	return (*fptr)(hSessionHandle, pucData,  uiDataLength);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_HashUpdate");
	return (*fptr)(hSessionHandle, pucData,  uiDataLength);
#endif
}
SGD_RV SDFHashFinal(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UCHAR_PRT *pucHash,SGD_UINT32  *puiHashLength)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UCHAR *,SGD_UINT32  *);
	*pucHash = calloc(*puiHashLength, sizeof(SGD_UCHAR));
	if (*pucHash == NULL) {
		return SGD_FALSE;
	}
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_HashFinal");
	return (*fptr)(hSessionHandle, *pucHash,  puiHashLength);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_HashFinal");
	return (*fptr)(hSessionHandle, *pucHash,  puiHashLength);
#endif
}



SGD_RV SDFGetSymmKeyHandle(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32 uiKeyIndex, SGD_HANDLE *phKeyHandle)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE , SGD_UINT32 , SGD_HANDLE *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_GetSymmKeyHandle");
	return (*fptr)(hSessionHandle,  uiKeyIndex,  phKeyHandle);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_GetSymmKeyHandle");
	return (*fptr)(hSessionHandle,  uiKeyIndex,  phKeyHandle);
#endif
}
SGD_RV SDFImportKey(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UCHAR_PRT pucKey, SGD_UINT32 uiKeyLength,SGD_HANDLE *phKeyHandle)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE , SGD_UCHAR *, SGD_UINT32 ,SGD_HANDLE *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ImportKey");
	return (*fptr)(hSessionHandle,  pucKey,  uiKeyLength, phKeyHandle);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ImportKey");
	return (*fptr)(hSessionHandle,  pucKey,  uiKeyLength, phKeyHandle);
#endif
}
SGD_RV SDFExternalPrivateKeyOperation_RSA(struct LibHandle * h,SGD_HANDLE hSessionHandle, RSArefPrivateKey *pucPrivateKey,SGD_UCHAR_PRT *pucDataInput,SGD_UINT32  uiInputLength,SGD_UCHAR_PRT *pucDataOutput,SGD_UINT32  *puiOutputLength)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE , RSArefPrivateKey *,SGD_UCHAR *,SGD_UINT32  ,SGD_UCHAR *,SGD_UINT32  *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ExternalPrivateKeyOperation_RSA");
	return (*fptr)(hSessionHandle,  pucPrivateKey, *pucDataInput,  uiInputLength, *pucDataOutput,  puiOutputLength);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ExternalPrivateKeyOperation_RSA");
	return (*fptr)(hSessionHandle,  pucPrivateKey, *pucDataInput,  uiInputLength, *pucDataOutput,  puiOutputLength);
#endif
}
SGD_RV SDFExternalSign_ECC(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UINT32 uiAlgID,ECCrefPrivateKey *pucPrivateKey,SGD_UCHAR_PRT pucData,SGD_UINT32  uiDataLength,ECCSignature *pucSignature)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UINT32 ,ECCrefPrivateKey *,SGD_UCHAR *,SGD_UINT32  ,ECCSignature *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ExternalSign_ECC");
	return (*fptr)(hSessionHandle, uiAlgID, pucPrivateKey, pucData,  uiDataLength, pucSignature);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ExternalSign_ECC");
	return (*fptr)(hSessionHandle, uiAlgID, pucPrivateKey, pucData,  uiDataLength, pucSignature);
#endif
}
SGD_RV SDFExternalDecrypt_ECC(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UINT32 uiAlgID,ECCrefPrivateKey *pucPrivateKey,ECCCipher *pucEncData,SGD_UCHAR_PRT *pucData,SGD_UINT32  *puiDataLength)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UINT32 ,ECCrefPrivateKey *,ECCCipher *,SGD_UCHAR *,SGD_UINT32  *);
	*pucData = calloc(*puiDataLength, sizeof(SGD_UCHAR));
	if (*pucData == NULL) {
		return SGD_FALSE;
	}
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ExternalDecrypt_ECC");
	return (*fptr)(hSessionHandle, uiAlgID, pucPrivateKey, pucEncData, *pucData,  puiDataLength);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ExternalDecrypt_ECC");
	return (*fptr)(hSessionHandle, uiAlgID, pucPrivateKey, pucEncData, *pucData,  puiDataLength);
#endif
}
SGD_RV SDFInternalDecrypt_ECC(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UINT32  uiISKIndex,SGD_UINT32 uiAlgID,ECCCipher *pucEncData,SGD_UCHAR_PRT *pucData,SGD_UINT32  *puiDataLength)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UINT32  ,SGD_UINT32 ,ECCCipher *,SGD_UCHAR *,SGD_UINT32  *);
	*pucData = calloc(*puiDataLength, sizeof(SGD_UCHAR));
	if (*pucData == NULL) {
		return SGD_FALSE;
	}
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_InternalDecrypt_ECC");
	return (*fptr)(hSessionHandle,  uiISKIndex, uiAlgID, pucEncData, *pucData,  puiDataLength);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_InternalDecrypt_ECC");
	return (*fptr)(hSessionHandle,  uiISKIndex, uiAlgID, pucEncData, *pucData,  puiDataLength);
#endif
}
SGD_RV SDFInternalEncrypt_ECC(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32  uiISKIndex, SGD_UINT32 uiAlgID, SGD_UCHAR_PRT pucData, SGD_UINT32  uiDataLength, ECCCipher *pucEncData)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE , SGD_UINT32  , SGD_UINT32 , SGD_UCHAR *, SGD_UINT32  , ECCCipher *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_InternalEncrypt_ECC");
	return (*fptr)(hSessionHandle,   uiISKIndex,  uiAlgID,  pucData,   uiDataLength,  pucEncData);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_InternalEncrypt_ECC");
	return (*fptr)(hSessionHandle,   uiISKIndex,  uiAlgID,  pucData,   uiDataLength,  pucEncData);
#endif
}


SGD_RV SDFExportKeyWithEPK_RSA(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_HANDLE hKeyHandle, RSArefPublicKey *pucPublicKey, SGD_UCHAR_PRT *pucKey, SGD_UINT32 *puiKeyLength)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE , SGD_HANDLE , RSArefPublicKey *, SGD_UCHAR *, SGD_UINT32 *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ExportKeyWithEPK_RSA");
	return (*fptr)(hSessionHandle,  hKeyHandle,  pucPublicKey,  *pucKey,  puiKeyLength);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ExportKeyWithEPK_RSA");
	return (*fptr)(hSessionHandle,  hKeyHandle,  pucPublicKey,  *pucKey,  puiKeyLength);
#endif
}
SGD_RV SDFExportKeyWithEPK_ECC(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_HANDLE hKeyHandle, SGD_UINT32 uiAlgID, ECCrefPublicKey *pucPublicKey, ECCCipher *pucKey)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE , SGD_HANDLE , SGD_UINT32 , ECCrefPublicKey *, ECCCipher *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ExportKeyWithEPK_ECC");
	return (*fptr)(hSessionHandle,  hKeyHandle,  uiAlgID,  pucPublicKey,  pucKey);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ExportKeyWithEPK_ECC");
	return (*fptr)(hSessionHandle,  hKeyHandle,  uiAlgID,  pucPublicKey,  pucKey);
#endif
}
SGD_RV SDFExportKeyWithKEK(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_HANDLE hKeyHandle, SGD_UINT32 uiAlgID, SGD_UINT32 uiKEKIndex, SGD_UCHAR_PRT *pucKey, SGD_UINT32 *puiKeyLength)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE , SGD_HANDLE , SGD_UINT32 , SGD_UINT32 , SGD_UCHAR *, SGD_UINT32 *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ExportKeyWithKEK");
	return (*fptr)(hSessionHandle,  hKeyHandle,  uiAlgID,  uiKEKIndex,  *pucKey,  puiKeyLength);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ExportKeyWithKEK");
	return (*fptr)(hSessionHandle,  hKeyHandle,  uiAlgID,  uiKEKIndex,  *pucKey,  puiKeyLength);
#endif
}


SGD_RV SDFExportSignMasterPublicKey_SM9(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UINT32  uiKeyIndex,SM9refSignMasterPublicKey *pPublicKey)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UINT32  ,SM9refSignMasterPublicKey *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ExportSignMasterPublicKey_SM9");
	return (*fptr)(hSessionHandle,  uiKeyIndex, pPublicKey);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ExportSignMasterPublicKey_SM9");
	return (*fptr)(hSessionHandle,  uiKeyIndex, pPublicKey);
#endif
}
SGD_RV SDFExportEncMasterPublicKey_SM9(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UINT32  uiKeyIndex,SM9refEncMasterPublicKey *pPublicKey)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UINT32  ,SM9refEncMasterPublicKey *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ExportEncMasterPublicKey_SM9");
	return (*fptr)(hSessionHandle, uiKeyIndex, pPublicKey);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ExportEncMasterPublicKey_SM9");
	return (*fptr)(hSessionHandle, uiKeyIndex, pPublicKey);
#endif
}
SGD_RV SDFExportSignMasterKeyPairG_SM9(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UINT32  uiKeyIndex,SGD_UCHAR_PRT *pPairG,SGD_UINT32 *puiPairGLen)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UINT32  ,SGD_UCHAR *,SGD_UINT32 *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ExportSignMasterKeyPairG_SM9");
	return (*fptr)(hSessionHandle,  uiKeyIndex, *pPairG, puiPairGLen);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ExportSignMasterKeyPairG_SM9");
	return (*fptr)(hSessionHandle,  uiKeyIndex, *pPairG, puiPairGLen);
#endif
}
SGD_RV SDFExportEncMasterKeyPairG_SM9(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UINT32  uiKeyIndex,SGD_UCHAR_PRT *pPairG,SGD_UINT32 *puiPairGLen)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UINT32  ,SGD_UCHAR *,SGD_UINT32 *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ExportEncMasterKeyPairG_SM9");
	return (*fptr)(hSessionHandle,  uiKeyIndex, *pPairG, puiPairGLen);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ExportEncMasterKeyPairG_SM9");
	return (*fptr)(hSessionHandle,  uiKeyIndex, *pPairG, puiPairGLen);
#endif
}
SGD_RV SDFImportUserSignPrivateKey_SM9(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UINT32 uiKeyIndex,SM9refSignUserPrivateKey  *pUserPrivateKey)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UINT32 ,SM9refSignUserPrivateKey  *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ImportUserSignPrivateKey_SM9");
	return (*fptr)(hSessionHandle, uiKeyIndex,  pUserPrivateKey);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ImportUserSignPrivateKey_SM9");
	return (*fptr)(hSessionHandle, uiKeyIndex,  pUserPrivateKey);
#endif
}
SGD_RV SDFImportUserEncPrivateKey_SM9(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32 uiKeyIndex, SM9refEncUserPrivateKey  *pUserPrivateKey)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE , SGD_UINT32 , SM9refEncUserPrivateKey  *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_ImportUserEncPrivateKey_SM9");
	return (*fptr)(hSessionHandle,  uiKeyIndex,   pUserPrivateKey);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_ImportUserEncPrivateKey_SM9");
	return (*fptr)(hSessionHandle,  uiKeyIndex,   pUserPrivateKey);
#endif
}
SGD_RV SDFGenerateSignUserPrivateKey_SM9(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32 uiKeyIndex, SGD_UCHAR hid, SGD_UCHAR_PRT *pucUserID, SGD_UINT32 uiUserIDLen, SM9refSignUserPrivateKey  *pUserPrivateKey)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE , SGD_UINT32 , SGD_UCHAR , SGD_UCHAR *, SGD_UINT32 , SM9refSignUserPrivateKey  *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_GenerateSignUserPrivateKey_SM9");
	return (*fptr)(hSessionHandle,  uiKeyIndex,  hid,  *pucUserID,  uiUserIDLen,   pUserPrivateKey);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_GenerateSignUserPrivateKey_SM9");
	return (*fptr)(hSessionHandle,  uiKeyIndex,  hid,  *pucUserID,  uiUserIDLen,   pUserPrivateKey);
#endif
}
SGD_RV SDFGenerateEncUserPrivateKey_SM9(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UINT32 uiKeyIndex, SGD_UCHAR hid,SGD_UCHAR_PRT *pucUserID,SGD_UINT32 uiUserIDLen,SM9refEncUserPrivateKey  *pUserPrivateKey)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE , SGD_UINT32 , SGD_UCHAR ,SGD_UCHAR *,SGD_UINT32 ,SM9refEncUserPrivateKey  *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_GenerateEncUserPrivateKey_SM9");
	return (*fptr)(hSessionHandle,  uiKeyIndex,  hid, *pucUserID, uiUserIDLen,  pUserPrivateKey);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_GenerateEncUserPrivateKey_SM9");
	return (*fptr)(hSessionHandle,  uiKeyIndex,  hid, *pucUserID, uiUserIDLen,  pUserPrivateKey);
#endif
}
SGD_RV SDFSign_SM9(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UINT32 uiKeyIndex,SM9refSignUserPrivateKey  *pUserPrivateKey,SM9refSignMasterPublicKey *pMasterPublicKey,SGD_UCHAR_PRT *pucDataInput,SGD_UINT32 uiDataInputLen,SM9Signature  *pSignature)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UINT32 ,SM9refSignUserPrivateKey  *,SM9refSignMasterPublicKey *,SGD_UCHAR *,SGD_UINT32 ,SM9Signature  *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_Sign_SM9");
	return (*fptr)(hSessionHandle, uiKeyIndex,  pUserPrivateKey, pMasterPublicKey, *pucDataInput, uiDataInputLen,  pSignature);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_Sign_SM9");
	return (*fptr)(hSessionHandle, uiKeyIndex,  pUserPrivateKey, pMasterPublicKey, *pucDataInput, uiDataInputLen,  pSignature);
#endif
}
SGD_RV SDFSignEx_SM9(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UINT32 uiKeyIndex,SM9refSignUserPrivateKey  *pUserPrivateKey,SM9refSignMasterPublicKey *pMasterPublicKey,SGD_UCHAR_PRT *pPairG,SGD_UINT32 uiPairGLen,SGD_UCHAR_PRT *pucDataInput,SGD_UINT32 uiDataInputLen,SM9Signature  *pSignature)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UINT32 ,SM9refSignUserPrivateKey  *,SM9refSignMasterPublicKey *,SGD_UCHAR *,SGD_UINT32 ,SGD_UCHAR *,SGD_UINT32 ,SM9Signature  *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_SignEx_SM9");
	return (*fptr)(hSessionHandle, uiKeyIndex,  pUserPrivateKey, pMasterPublicKey, *pPairG, uiPairGLen, *pucDataInput, uiDataInputLen,  pSignature);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_SignEx_SM9");
	return (*fptr)(hSessionHandle, uiKeyIndex,  pUserPrivateKey, pMasterPublicKey, *pPairG, uiPairGLen, *pucDataInput, uiDataInputLen,  pSignature);
#endif
}
SGD_RV SDFVerify_SM9(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UCHAR hid,SGD_UCHAR_PRT *pucUserID,SGD_UINT32  uiUserIDLen,SM9refSignMasterPublicKey  *pMasterPublicKey,SGD_UCHAR_PRT *pucData,SGD_UINT32   uiDataInputLen,SM9Signature  *pSignature)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UCHAR ,SGD_UCHAR *,SGD_UINT32  ,SM9refSignMasterPublicKey  *,SGD_UCHAR *,SGD_UINT32   ,SM9Signature  *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_Verify_SM9");
	return (*fptr)(hSessionHandle, hid, *pucUserID,  uiUserIDLen,  pMasterPublicKey, *pucData,   uiDataInputLen,  pSignature);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_Verify_SM9");
	return (*fptr)(hSessionHandle, hid, *pucUserID,  uiUserIDLen,  pMasterPublicKey, *pucData,   uiDataInputLen,  pSignature);
#endif
}
SGD_RV SDFVerifyEx_SM9(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UCHAR hid,SGD_UCHAR_PRT *pucUserID,SGD_UINT32 uiUserIDLen,SM9refSignMasterPublicKey  *pMasterPublicKey,SGD_UCHAR_PRT *pPairG,SGD_UINT32 uiPairGLen,SGD_UCHAR_PRT *pucData,SGD_UINT32   uiDataInputLen,SM9Signature  *pSignature)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE hSessionHandle,SGD_UCHAR hid,SGD_UCHAR *pucUserID,SGD_UINT32 uiUserIDLen,SM9refSignMasterPublicKey  *pMasterPublicKey,SGD_UCHAR *pPairG,SGD_UINT32 uiPairGLen,SGD_UCHAR *pucData,SGD_UINT32   uiDataInputLen,SM9Signature  *pSignature);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_VerifyEx_SM9");
	return (*fptr)(hSessionHandle, hid, *pucUserID, uiUserIDLen,  pMasterPublicKey, *pPairG, uiPairGLen, *pucData,   uiDataInputLen,  pSignature);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_VerifyEx_SM9");
	return (*fptr)(hSessionHandle, hid, *pucUserID, uiUserIDLen,  pMasterPublicKey, *pPairG, uiPairGLen, *pucData,   uiDataInputLen,  pSignature);
#endif
}
SGD_RV SDFEncrypt_SM9(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UCHAR hid,SGD_UCHAR_PRT *pucUserID,SGD_UINT32  uiUserIDLen,SM9refEncMasterPublicKey *pPubluicKey,SGD_UCHAR_PRT *pucData,SGD_UINT32   uiDataLength,SM9Cipher *pCipher)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UCHAR ,SGD_UCHAR *,SGD_UINT32  ,SM9refEncMasterPublicKey *,SGD_UCHAR *,SGD_UINT32   ,SM9Cipher *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_Encrypt_SM9");
	return (*fptr)(hSessionHandle, hid, *pucUserID,  uiUserIDLen, pPubluicKey, *pucData,   uiDataLength, pCipher);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_Encrypt_SM9");
	return (*fptr)(hSessionHandle, hid, *pucUserID,  uiUserIDLen, pPubluicKey, *pucData,   uiDataLength, pCipher);
#endif
}
SGD_RV SDFEncryptEx_SM9(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UCHAR hid,SGD_UCHAR_PRT *pucUserID,SGD_UINT32  uiUserIDLen,SM9refEncMasterPublicKey *pPubluicKey,SGD_UCHAR_PRT *pPairG,SGD_UINT32  nPairGLen,SGD_UCHAR_PRT *pucData,SGD_UINT32   uiDataLength,SM9Cipher *pCipher)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UCHAR ,SGD_UCHAR *,SGD_UINT32  ,SM9refEncMasterPublicKey *,SGD_UCHAR *,SGD_UINT32  ,SGD_UCHAR *,SGD_UINT32   ,SM9Cipher *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_EncryptEx_SM9");
	return (*fptr)(hSessionHandle, hid, *pucUserID,  uiUserIDLen, pPubluicKey, *pPairG,  nPairGLen, *pucData,   uiDataLength, pCipher);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_EncryptEx_SM9");
	return (*fptr)(hSessionHandle, hid, *pucUserID,  uiUserIDLen, pPubluicKey, *pPairG,  nPairGLen, *pucData,   uiDataLength, pCipher);
#endif
}
SGD_RV SDFDecrypt_SM9(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UCHAR_PRT *pucUserID,SGD_UINT32  uiUserIDLen,SGD_UINT32 uiKeyIndex,SM9refEncUserPrivateKey  *pUserPrivateKey,SM9Cipher * pCipher,SGD_UCHAR_PRT *pucPlainData,SGD_UINT32  *uiPlainDataLength)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UCHAR *,SGD_UINT32  ,SGD_UINT32 ,SM9refEncUserPrivateKey  *,SM9Cipher * ,SGD_UCHAR *,SGD_UINT32  *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_Decrypt_SM9");
	return (*fptr)(hSessionHandle, *pucUserID,  uiUserIDLen, uiKeyIndex,  pUserPrivateKey,  pCipher, *pucPlainData,  uiPlainDataLength);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_Decrypt_SM9");
	return (*fptr)(hSessionHandle, *pucUserID,  uiUserIDLen, uiKeyIndex,  pUserPrivateKey,  pCipher, *pucPlainData,  uiPlainDataLength);
#endif
}
SGD_RV SDFEncap_SM9(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UCHAR hid,SGD_UCHAR_PRT *pucUserID,SGD_UINT32  uiUserIDLen,SM9refEncMasterPublicKey  *pPublicKey,SGD_UINT32 uiKeyLen,SGD_UCHAR_PRT *pKey,SM9refKeyPackage *pKeyPackage)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UCHAR ,SGD_UCHAR *,SGD_UINT32  ,SM9refEncMasterPublicKey  *,SGD_UINT32 ,SGD_UCHAR *,SM9refKeyPackage *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_Encap_SM9");
	return (*fptr)(hSessionHandle, hid, *pucUserID,  uiUserIDLen,  pPublicKey, uiKeyLen, *pKey, pKeyPackage);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_Encap_SM9");
	return (*fptr)(hSessionHandle, hid, *pucUserID,  uiUserIDLen,  pPublicKey, uiKeyLen, *pKey, pKeyPackage);
#endif
}
SGD_RV SDFDecap_SM9(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UCHAR_PRT *pucUserID,SGD_UINT32  uiUserIDLen,SGD_UINT32 uiKeyIndex,SM9refEncUserPrivateKey  *pUserPrivateKey,SM9refKeyPackage *pKeyPackage,SGD_UINT32  uiKeyLen,SGD_UCHAR_PRT *pucKey)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UCHAR *,SGD_UINT32  ,SGD_UINT32 ,SM9refEncUserPrivateKey  *,SM9refKeyPackage *,SGD_UINT32  ,SGD_UCHAR *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_Decap_SM9");
	return (*fptr)(hSessionHandle, *pucUserID,  uiUserIDLen, uiKeyIndex,  pUserPrivateKey, pKeyPackage,  uiKeyLen, *pucKey);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_Decap_SM9");
	return (*fptr)(hSessionHandle, *pucUserID,  uiUserIDLen, uiKeyIndex,  pUserPrivateKey, pKeyPackage,  uiKeyLen, *pucKey);
#endif
}
SGD_RV SDFGenerateAgreementDataWithSM9(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UCHAR hid, SGD_UCHAR_PRT *pucResponseID, SGD_UINT32 uiResponseIDLength, SM9refEncMasterPublicKey  *pPublicKey, SM9refEncMasterPublicKey  *pucSponsorTmpPublicKey, SGD_HANDLE *phAgreementHandle)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE , SGD_UCHAR , SGD_UCHAR *, SGD_UINT32 , SM9refEncMasterPublicKey  *, SM9refEncMasterPublicKey  *, SGD_HANDLE *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_GenerateAgreementDataWithSM9");
	return (*fptr)(hSessionHandle,  hid,  *pucResponseID,  uiResponseIDLength,   pPublicKey,   pucSponsorTmpPublicKey,  phAgreementHandle);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_GenerateAgreementDataWithSM9");
	return (*fptr)(hSessionHandle,  hid,  *pucResponseID,  uiResponseIDLength,   pPublicKey,   pucSponsorTmpPublicKey,  phAgreementHandle);
#endif
}
SGD_RV SDFGenerateAgreemetDataAndKeyWithSM9(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UINT32 uiKeyLen,SGD_UCHAR hid,SGD_UCHAR_PRT * pucResponseID,SGD_UINT32 uiResponseIDLen,SGD_UCHAR_PRT * pucSponsorID,SGD_UINT32 uiSponsorIDLen,SGD_UINT32 uiKeyIndex,SM9refEncUserPrivateKey  *pucResponsePrivateKey,SM9refEncMasterPublicKey *pucPublicKey,SM9refEncMasterPublicKey * pucSponsorTmpPublicKey,SM9refEncMasterPublicKey * pucResponseTmpPublicKey,SGD_UCHAR_PRT *pucHashSB,SGD_UINT32 *puiSBLen,SGD_UCHAR_PRT  *pucHashS2,SGD_UINT32 *puiS2Len,SGD_HANDLE *phKeyHandle)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UINT32 ,SGD_UCHAR ,SGD_UCHAR * ,SGD_UINT32 ,SGD_UCHAR * ,SGD_UINT32 ,SGD_UINT32 ,SM9refEncUserPrivateKey  *,SM9refEncMasterPublicKey *,SM9refEncMasterPublicKey * ,SM9refEncMasterPublicKey * ,SGD_UCHAR *,SGD_UINT32 *,SGD_UCHAR  *,SGD_UINT32 *,SGD_HANDLE *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_GenerateAgreemetDataAndKeyWithSM9");
	return (*fptr)(hSessionHandle, uiKeyLen, hid, *pucResponseID, uiResponseIDLen,  *pucSponsorID, uiSponsorIDLen, uiKeyIndex,  pucResponsePrivateKey, pucPublicKey,  pucSponsorTmpPublicKey,  pucResponseTmpPublicKey, *pucHashSB, puiSBLen,  *pucHashS2, puiS2Len, phKeyHandle);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_GenerateAgreemetDataAndKeyWithSM9");
	return (*fptr)(hSessionHandle, uiKeyLen, hid, *pucResponseID, uiResponseIDLen,  *pucSponsorID, uiSponsorIDLen, uiKeyIndex,  pucResponsePrivateKey, pucPublicKey,  pucSponsorTmpPublicKey,  pucResponseTmpPublicKey, *pucHashSB, puiSBLen,  *pucHashS2, puiS2Len, phKeyHandle);
#endif
}
SGD_RV SDFGenerateKeyWithSM9(struct LibHandle * h,SGD_HANDLE hSessionHandle,SGD_UINT32 uiKeyLen,SGD_UCHAR hid,SGD_UCHAR_PRT *pucSponsorID,SGD_UINT32 uiSponsorIDLen,SGD_UCHAR_PRT *pucResponseID,SGD_UINT32 uiResponseIDLen,SGD_UINT32 uiKeyIndex,SM9refEncUserPrivateKey   *pucSponsorPrivateKey,SM9refEncMasterPublicKey   *pucPublicKey,SM9refEncMasterPublicKey   *pucResponseTmpPublicKey,SGD_UCHAR_PRT *pucHashSB,SGD_UINT32 uiSBLen,SGD_UCHAR_PRT *pucHashSA,SGD_UINT32 *puiSALen,SGD_HANDLE hAgreementHandle,SGD_HANDLE *phKeyHandle)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE ,SGD_UINT32 ,SGD_UCHAR ,SGD_UCHAR *,SGD_UINT32 ,SGD_UCHAR *,SGD_UINT32 ,SGD_UINT32 ,SM9refEncUserPrivateKey   *,SM9refEncMasterPublicKey   *,SM9refEncMasterPublicKey   *,SGD_UCHAR *,SGD_UINT32 ,SGD_UCHAR *,SGD_UINT32 *,SGD_HANDLE ,SGD_HANDLE *);
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_GenerateKeyWithSM9");
	return (*fptr)(hSessionHandle, uiKeyLen, hid, *pucSponsorID, uiSponsorIDLen, *pucResponseID, uiResponseIDLen, uiKeyIndex,   pucSponsorPrivateKey,   pucPublicKey,   pucResponseTmpPublicKey, *pucHashSB, uiSBLen, *pucHashSA, puiSALen, hAgreementHandle, phKeyHandle);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_GenerateKeyWithSM9");
	return (*fptr)(hSessionHandle, uiKeyLen, hid, *pucSponsorID, uiSponsorIDLen, *pucResponseID, uiResponseIDLen, uiKeyIndex,   pucSponsorPrivateKey,   pucPublicKey,   pucResponseTmpPublicKey, *pucHashSB, uiSBLen, *pucHashSA, puiSALen, hAgreementHandle, phKeyHandle);
#endif
}
SGD_RV SDFGenerateKeyVerifySM9(struct LibHandle * h,SGD_HANDLE hSessionHandle, SGD_UCHAR_PRT *pHashS2, SGD_UINT32  uiS2Len, SGD_UCHAR_PRT *pHashSA, SGD_UINT32 uiSALen)
{
    typedef SGD_RV (*FPTR)(SGD_HANDLE , SGD_UCHAR *, SGD_UINT32  , SGD_UCHAR *, SGD_UINT32 );
#ifdef _WIN32
	FPTR fptr = (FPTR)GetProcAddress(h->handle, "SDF_GenerateKeyVerifySM9");
	return (*fptr)(hSessionHandle,  *pHashS2,   uiS2Len,  *pHashSA,  uiSALen);
#else
	FPTR fptr = (FPTR)dlsym(h->handle, "SDF_GenerateKeyVerifySM9");
	return (*fptr)(hSessionHandle,  *pHashS2,   uiS2Len,  *pHashSA,  uiSALen);
#endif
}
*/
import "C"
import (
	"fmt"
	"github.com/yzwskyspace/sdf/core"
	"os"
	"strings"
	"unsafe"
)

func ConvertToDeviceInfoGo(deviceInfo1 C.DEVICEINFO)(deviceInfo core.DeviceInfo){
	deviceInfo = core.DeviceInfo{
		IssuerName:  strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&deviceInfo1.IssuerName[0]), 40)), " "),
		DeviceName:  strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&deviceInfo1.DeviceName[0]), 16)), " "),
		DeviceSerial:  strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&deviceInfo1.DeviceSerial[0]), 16)), " "),
		DeviceVersion: uint(deviceInfo1.DeviceVersion),
		StandardVersion: uint(deviceInfo1.StandardVersion),
		SymAlgAbility: uint(deviceInfo1.SymAlgAbility),
		HashAlgAbility: uint(deviceInfo1.HashAlgAbility),
		BufferSize: uint(deviceInfo1.BufferSize),
	}
	temp1:=C.GoBytes(unsafe.Pointer(&deviceInfo1.AsymAlgAbility[0]),2)
	temp2:=C.GoBytes(unsafe.Pointer(&deviceInfo1.AsymAlgAbility[1]),2)
	deviceInfo.AsymAlgAbility[0]=uint(temp1[0])
	deviceInfo.AsymAlgAbility[1]=uint(temp2[0])
	return deviceInfo
}
func ConvertToRSArefPrivateKeyGo(pucPrivateKey C.RSArefPrivateKey)(privateKey core.RSArefPrivateKey){
	privateKey =core.RSArefPrivateKey{
		Bits: uint(pucPrivateKey.bits),
		M: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucPrivateKey.m[0]), 256)), " "),
		E: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucPrivateKey.e[0]), 256)), " "),
		D: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucPrivateKey.d[0]), 256)), " "),
		Coef: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucPrivateKey.coef[0]), 128)), " "),
	}
	privateKey.Prime[0] = strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucPrivateKey.prime[0]), 128)), " ")
	privateKey.Prime[1] = strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucPrivateKey.prime[1]), 128)), " ")
	privateKey.Pexp[0] = strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucPrivateKey.pexp[0]), 128)), " ")
	privateKey.Pexp[1] = strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucPrivateKey.pexp[1]), 128)), " ")
	return privateKey
}

func ConvertToRSArefPublicKeyGo(pucPublicKey C.RSArefPublicKey)(publicKey core.RSArefPublicKey){
	publicKey =core.RSArefPublicKey{
		Bits: uint(pucPublicKey.bits),
		M: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucPublicKey.m[0]), 256)), " "),
		E: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucPublicKey.e[0]), 256)), " "),
	}
	return publicKey
}



func  ConvertToECCrefPublicKeyGo(pucPublicKey C.ECCrefPublicKey)(publicKey core.ECCrefPublicKey){
	publicKey =core.ECCrefPublicKey{
		Bits: uint(pucPublicKey.bits),
		X: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucPublicKey.x[0]), 32)), " "),
		Y: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucPublicKey.y[0]), 32)), " "),
	}
	return publicKey
}

func ConvertToECCrefPrivateKeyGo(pucPrivateKey C.ECCrefPrivateKey)(privateKey core.ECCrefPrivateKey){
	privateKey =core.ECCrefPrivateKey{
		Bits: uint(pucPrivateKey.bits),
		K: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucPrivateKey.K[0]), 32)), " "),
	}
	return privateKey
}

func ConvertToECCCipherGo(pucKey C.ECCCipher)(key core.ECCCipher){
	key =core.ECCCipher{
		X: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucKey.x[0]), 64)), " "),
		Y: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucKey.y[0]), 64)), " "),
		M: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucKey.M[0]), 32)), " "),
		L: uint(pucKey.L),
		C: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucKey.C[0]), C.int(pucKey.L))), " "),
	}

	return key
}

func ConvertToECCSignatureGo(pucSignature C.ECCSignature)(signature core.ECCSignature){
	signature =core.ECCSignature{
		R: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucSignature.r[0]), 64)), " "),
		S: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pucSignature.s[0]), 64)), " "),
	}
	return signature
}

func ConvertToSM9refSignMasterPublicKeyGo(pSignMasterPublicKey C.SM9refSignMasterPublicKey)(signMasterPublicKey core.SM9refSignMasterPublicKey){
	signMasterPublicKey =core.SM9refSignMasterPublicKey{
		Bits: uint(pSignMasterPublicKey.bits),
		Xa: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pSignMasterPublicKey.xa[0]), 256)), " "),
		Xb: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pSignMasterPublicKey.xb[0]), 256)), " "),
		Ya: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pSignMasterPublicKey.ya[0]), 256)), " "),
		Yb: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pSignMasterPublicKey.yb[0]), 256)), " "),
	}
	return signMasterPublicKey
}

func ConvertToSM9refEncMasterPublicKeyGo(pEncMasterPublicKey C.SM9refEncMasterPublicKey)(encMasterPublicKey core.SM9refEncMasterPublicKey){
	encMasterPublicKey =core.SM9refEncMasterPublicKey{
		Bits: uint(pEncMasterPublicKey.bits),
		X: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pEncMasterPublicKey.x[0]), 256)), " "),
		Y: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pEncMasterPublicKey.y[0]), 256)), " "),
	}
	return encMasterPublicKey
}

func ConvertToSM9refEncUserPrivateKeyGo(pEncUserPrivateKey C.SM9refEncUserPrivateKey)(encUserPrivateKey core.SM9refEncUserPrivateKey){
	encUserPrivateKey = core.SM9refEncUserPrivateKey{
		Bits: uint(pEncUserPrivateKey.bits),
		Xa: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pEncUserPrivateKey.xa[0]), 256)), " "),
		Xb: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pEncUserPrivateKey.xb[0]), 256)), " "),
		Ya: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pEncUserPrivateKey.ya[0]), 256)), " "),
		Yb: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pEncUserPrivateKey.yb[0]), 256)), " "),
	}
	return encUserPrivateKey
}

func ConvertToSM9refSignUserPrivateKeyGo(pSignUserPrivateKey C.SM9refSignUserPrivateKey)(signUserPrivateKey core.SM9refSignUserPrivateKey){
	signUserPrivateKey=core.SM9refSignUserPrivateKey{
		Bits: uint(pSignUserPrivateKey.bits),
		X: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pSignUserPrivateKey.x[0]), 256)), " "),
		Y: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pSignUserPrivateKey.y[0]), 256)), " "),
	}
	return signUserPrivateKey
}
func ConvertToSM9CipherGo(pCipher C.SM9Cipher)(cipher core.SM9Cipher){
	cipher=core.SM9Cipher{
		X: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pCipher.x[0]), 256)), " "),
		Y: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pCipher.y[0]), 256)), " "),
		H: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pCipher.h[0]), 256)), " "),
		L: uint(pCipher.L),
		C: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pCipher.C[0]), 256)), " "),
	}
	return cipher
}

func ConvertToSM9SignatureGo(pSignature C.SM9Signature)(signature core.SM9Signature){
	signature = core.SM9Signature{
		H: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pSignature.h[0]), 256)), " "),
		X: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pSignature.x[0]), 256)), " "),
		Y: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pSignature.y[0]), 256)), " "),
	}
	return signature
}

func ConvertToSM9refKeyPackageGo(pKeyPackage C.SM9refKeyPackage)(keyPackage core.SM9refKeyPackage){
	keyPackage=core.SM9refKeyPackage{
		X: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pKeyPackage.x[0]), 256)), " "),
		Y: strings.TrimRight(string(C.GoBytes(unsafe.Pointer(&pKeyPackage.y[0]), 256)), " "),
	}
	return keyPackage
}

func New(libPath string) *Ctx{
	if x:=os.Getenv("SDFHSM_CONF");x==""{
		os.Setenv("SDFHSM_CONF",libPath)
	}else {
		libPath = x
	}
	c := new(Ctx)
	mod := C.CString(libPath)
	defer C.free(unsafe.Pointer(mod))
	c.libHandle = C.New(mod)
	if c.libHandle == nil{
		return nil
	}
	return  c
}
type  Error uint


func (e Error) Error() string{
	return fmt.Sprintf("sdf: 0x%X:%s",uint(e),core.StrErrors[uint(e)])
}

func ToError(e C.SGD_RV)error{
	if e == C.SDR_OK{
		return nil
	}
	return Error(e)
}

func deepCopy(src []byte)(dst []byte){
	dst=make([]byte, len(src))
	for i,v:=range src{
		dst[i]=v
	}
	return
}

type Ctx struct {
	libHandle     *C.struct_LibHandle
}

type DeviceHandleType   C.SGD_HANDLE
type SessionHandleType  C.SGD_HANDLE
type KeyHandleType    C.SGD_HANDLE
type AgreementHandleType C.SGD_HANDLE

var stubData = []byte{0}
func CMessage(data []byte) (dataPtr C.SGD_UCHAR_PRT) {
	l := len(data)
	if l == 0 {
		data = stubData
	}
	dataPtr=C.SGD_UCHAR_PRT(unsafe.Pointer(&data[0]))
	return dataPtr
}


//1.打开设备
func (c *Ctx)SDFOpenDevice(deviceHandle  DeviceHandleType) (deviceHandle2 DeviceHandleType,err error){
    var err1 C.SGD_RV
    var dH =C.SGD_HANDLE(deviceHandle)
	err1 = C.SDFOpenDevice(c.libHandle,&dH)
	err = ToError(err1)

	deviceHandle2 = DeviceHandleType(dH)
	if err!= nil {
		return nil, err
	}
	return deviceHandle2,err
}
//2.关闭设备
func (c *Ctx)SDFCloseDevice(deviceHandle  DeviceHandleType)(err error){
	var err1 C.SGD_RV
	err1 = C.SDFCloseDevice(c.libHandle,C.SGD_HANDLE(deviceHandle))
	err = ToError(err1)
	return err
}
//3.创建会话
func (c *Ctx)SDFOpenSession(deviceHandle  DeviceHandleType)  (sessionHandle SessionHandleType,err error){
	var err1 C.SGD_RV
	var s C.SGD_HANDLE
	err1 = C.SDFOpenSession(c.libHandle,C.SGD_HANDLE(deviceHandle),&s)
	sessionHandle=SessionHandleType(s)
	err=ToError(err1)
	return sessionHandle,err
}
//4.关闭会话
func (c *Ctx)SDFCloseSession(sessionHandle  SessionHandleType)(err error){
	var err1 C.SGD_RV
	err1 = C.SDFCloseSession(c.libHandle,C.SGD_HANDLE(sessionHandle))
	err = ToError(err1)
	return nil
}
//5.获取设备信息
func (c *Ctx)SDFGetDeviceInfo(sessionHandle  SessionHandleType) (deviceInfo core.DeviceInfo,err error){
	var deviceInfo1 C.DEVICEINFO
	var err1 C.SGD_RV
	err1 = C.SDFGetDeviceInfo(c.libHandle,C.SGD_HANDLE(sessionHandle),&deviceInfo1)
	deviceInfo =ConvertToDeviceInfoGo(deviceInfo1)
	err= ToError(err1)
	return deviceInfo,err
}
//6.产生随机数
func (c *Ctx)SDFGenerateRandom(sessionHandle SessionHandleType,length uint ) (randomData []byte,err error){
	var err1 C.SGD_RV
	var random C.SGD_UCHAR_PRT
	err1 = C.SDFGenerateRandom(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(length),&random)
	err = ToError(err1)
	randomData = C.GoBytes(unsafe.Pointer(random), C.int(length))
	//C.free(unsafe.Pointer(random))
	return randomData,err
}
//7.获取私钥使用权限
func (c *Ctx)SDFGetPrivateKeyAccessRight(sessionHandle SessionHandleType,keyIndex uint,password []byte,pwdLength uint)(err error){
	var err1 C.SGD_RV
    err1 = C.SDFGetPrivateKeyAccessRight(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(keyIndex),CMessage(password),C.SGD_UINT32(pwdLength))
	err = ToError(err1)
	return err
}
//8.释放私钥使用权限
func (c *Ctx)SDFReleasePrivateKeyAccessRight(sessionHandle SessionHandleType,keyIndex uint) (err error){
	var err1 C.SGD_RV
	err1 = C.SDFReleasePrivateKeyAccessRight(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(keyIndex))
	err = ToError(err1)
	return err
}
//9.导出 ＲＳＡ 签名公钥
func (c *Ctx)SDFExportSignPublicKey_RSA(sessionHandle SessionHandleType,keyIndex uint)(publicKey core.RSArefPublicKey,err error){
	var err1 C.SGD_RV
	var pucPublicKey C.RSArefPublicKey
	err1 = C.SDFExportSignPublicKey_RSA(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(keyIndex),&pucPublicKey)
	publicKey = ConvertToRSArefPublicKeyGo(pucPublicKey)
	err = ToError(err1)
	return publicKey,err
}
//10.导出 ＲＳＡ 加密公钥
func (c *Ctx)SDFExportEncPublicKey_RSA(sessionHandle SessionHandleType,keyIndex uint)(publicKey core.RSArefPublicKey,err error){
	var err1 C.SGD_RV
	var pucPublicKey C.RSArefPublicKey
	err1 = C.SDFExportEncPublicKey_RSA(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(keyIndex),&pucPublicKey)
	publicKey = ConvertToRSArefPublicKeyGo(pucPublicKey)
	err = ToError(err1)
	return publicKey,err
}
//11.产生 ＲＳＡ 非对称密钥对并输出
func (c *Ctx)SDFGenerateKeyPair_RSA(sessionHandle SessionHandleType,uiKeyBits uint)(publicKey core.RSArefPublicKey,privateKey core.RSArefPrivateKey,err error){

	var err1 C.SGD_RV
	var pucPublicKey C.RSArefPublicKey
	var pucPrivateKey C.RSArefPrivateKey
	err1 = C.SDFGenerateKeyPair_RSA(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyBits),&pucPublicKey,&pucPrivateKey)
	publicKey = ConvertToRSArefPublicKeyGo(pucPublicKey)
	privateKey = ConvertToRSArefPrivateKeyGo(pucPrivateKey)
	err = ToError(err1)
	return publicKey,privateKey,err
}
//12.生成会话密钥并用内部 ＲＳＡ 公钥加密输出
func (c *Ctx)SDFGenerateKeyWithIPK_RSA(sessionHandle SessionHandleType,uiIPKIndex uint,uiKeyBits uint)(key []byte,keyLength uint,keyHandle KeyHandleType,err error){
	var err1 C.SGD_RV
	var pucKey C.SGD_UCHAR_PRT
	var length C.SGD_UINT32
	var phKeyHandle C.SGD_HANDLE
	err1 = C.SDFGenerateKeyWithIPK_RSA(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiIPKIndex),C.SGD_UINT32(uiKeyBits),&pucKey,&length,&phKeyHandle)
	key = C.GoBytes(unsafe.Pointer(pucKey), C.int(length))
	//C.free(unsafe.Pointer(pucKey))
	keyLength = uint(length)
	keyHandle=KeyHandleType(phKeyHandle)
	err=ToError(err1)
	return key,keyLength,keyHandle,err
}
//13.生成会话密钥并用外部 ＲＳＡ 公钥加密输出
func (c *Ctx)SDFGenerateKeyWithEPK_RSA(sessionHandle SessionHandleType,uiKeyBits uint,publicKey core.RSArefPublicKey)(key []byte,keyLength uint,keyHandle KeyHandleType,err error){
	var err1 C.SGD_RV
	var pucKey C.SGD_UCHAR_PRT
	var puiKeyLength C.SGD_UINT32
	var phKeyHandle C.SGD_HANDLE
	var pubKey C.RSArefPublicKey
	pubKey.bits = C.SGD_UINT32(publicKey.Bits)

	for i:=0;i<len(publicKey.M);i++{
		pubKey.m[i]=C.SGD_UCHAR(publicKey.M[i])
	}
	for i:=0;i<len(publicKey.E);i++{
		pubKey.e[i]=C.SGD_UCHAR(publicKey.E[i])
	}

	err1 = C.SDFGenerateKeyWithEPK_RSA(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyBits),&pubKey,&pucKey,&puiKeyLength,&phKeyHandle)
	key = C.GoBytes(unsafe.Pointer(pucKey), C.int(puiKeyLength))
	keyLength = uint(puiKeyLength)
	//C.free(unsafe.Pointer(pucKey))
	keyHandle=KeyHandleType(phKeyHandle)
	err=ToError(err1)
	return key,keyLength,keyHandle,err
}
//14.导入会话密钥并用内部 ＲＳＡ 私钥解密
func (c *Ctx)SDFImportKeyWithISK_RSA(sessionHandle SessionHandleType,uiKeyBits uint,uiKeyLength uint)([]byte,KeyHandleType,error){
	var err C.SGD_RV
	var pucKey C.SGD_UCHAR_PRT
	var phKeyHandle C.SGD_HANDLE
	err = C.SDFImportKeyWithISK_RSA(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyBits),&pucKey,C.SGD_UINT32(uiKeyLength),&phKeyHandle)
	p := C.GoBytes(unsafe.Pointer(pucKey), C.int(uiKeyLength))
	C.free(unsafe.Pointer(pucKey))
	return p,KeyHandleType(phKeyHandle),ToError(err)
}
//15.基于 ＲＳＡ 算法的数字信封转换
func (c *Ctx)SDFExchangeDigitEnvelopeBaseOnRSA(sessionHandle SessionHandleType,uiKeyIndex uint,uiDELength uint)(core.RSArefPublicKey,[]byte,[]byte,error){
	var err C.SGD_RV
	var pucPublicKey C.RSArefPublicKey
	var pucDEInput C.SGD_UCHAR_PRT
	var pucDEOutput C.SGD_UCHAR_PRT
	var puiDELength C.SGD_UINT32
	err = C.SDFExchangeDigitEnvelopeBaseOnRSA(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyIndex),&pucPublicKey,&pucDEInput,C.SGD_UINT32(uiDELength),&pucDEOutput,&puiDELength)
	publicKey := ConvertToRSArefPublicKeyGo(pucPublicKey)
	input := C.GoBytes(unsafe.Pointer(pucDEInput), C.int(uiDELength))
	C.free(unsafe.Pointer(pucDEInput))
	output := C.GoBytes(unsafe.Pointer(pucDEOutput), C.int(puiDELength))
	C.free(unsafe.Pointer(pucDEOutput))
	return publicKey,input,output,ToError(err)
}
//16.导出 ＥＣＣ签名公钥
func (c *Ctx)SDFExportSignPublicKey_ECC(sessionHandle SessionHandleType,uiKeyIndex uint)(publicKey core.ECCrefPublicKey,err error){
	var err1 C.SGD_RV
	var pucPublicKey C.ECCrefPublicKey
	err1 = C.SDFExportSignPublicKey_ECC(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyIndex),&pucPublicKey)
	publicKey =ConvertToECCrefPublicKeyGo(pucPublicKey)
	err = ToError(err1)
	return publicKey,err
}
//17.导出 ＥＣＣ加密公钥
func (c *Ctx)SDFExportEncPublicKey_ECC(sessionHandle SessionHandleType,uiKeyIndex uint)(publicKey core.ECCrefPublicKey,err error){
	var err1 C.SGD_RV
	var pucPublicKey C.ECCrefPublicKey
	err1 = C.SDFExportEncPublicKey_ECC(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyIndex),&pucPublicKey)
	publicKey =ConvertToECCrefPublicKeyGo(pucPublicKey)
	err = ToError(err1)
	return publicKey,err
}
//18.产生 ＥＣＣ非对称密钥对并输出
func (c *Ctx)SDFGenerateKeyPair_ECC(sessionHandle SessionHandleType,uiAlgID uint,uiKeyBits uint)(publicKey core.ECCrefPublicKey,privateKey core.ECCrefPrivateKey,err error){
	var err1 C.SGD_RV
	var pucPublicKey C.ECCrefPublicKey
	var pucPrivateKey C.ECCrefPrivateKey
	err1 = C.SDFGenerateKeyPair_ECC(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiAlgID),C.SGD_UINT32(uiKeyBits),&pucPublicKey,&pucPrivateKey)
	publicKey = ConvertToECCrefPublicKeyGo(pucPublicKey)
	privateKey = ConvertToECCrefPrivateKeyGo(pucPrivateKey)
	err = ToError(err1)
	return publicKey,privateKey,err
}
//19.生成会话密钥并用内部 ＥＣＣ公钥加密输出
func (c *Ctx)SDFGenerateKeyWithIPK_ECC(sessionHandle SessionHandleType,uiIPKIndex uint,uiKeyBits uint)(key core.ECCCipher,keyHandle KeyHandleType,err error){
	var err1 C.SGD_RV
	var pucKey C.ECCCipher
	var phKeyHandle C.SGD_HANDLE
	err1 = C.SDFGenerateKeyWithIPK_ECC(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiIPKIndex),C.SGD_UINT32(uiKeyBits),&pucKey,&phKeyHandle)
	key = ConvertToECCCipherGo(pucKey)
	keyHandle = KeyHandleType(phKeyHandle)
	err = ToError(err1)
	return key,keyHandle,err
}
//20.生成会话密钥并用外部 ＥＣＣ公钥加密输出
func (c *Ctx)SDFGenerateKeyWithEPK_ECC(sessionHandle SessionHandleType,uiKeyBits uint,uiAlgID uint,publicKey core.ECCrefPublicKey)(key core.ECCCipher,keyHandle KeyHandleType,err error){
	var err1 C.SGD_RV
	var pucPublicKey C.ECCrefPublicKey
	pucPublicKey.bits = C.SGD_UINT32(publicKey.Bits)

	for i:=0;i<len(publicKey.X);i++{
		pucPublicKey.x[i]=C.SGD_UCHAR(publicKey.Y[i])
	}
	for i:=0;i<len(publicKey.Y);i++{
		pucPublicKey.y[i]=C.SGD_UCHAR(publicKey.Y[i])
	}
	var pucKey C.ECCCipher
	var phKeyHandle C.SGD_HANDLE
	err1 = C.SDFGenerateKeyWithEPK_ECC(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyBits),C.SGD_UINT32(uiAlgID),&pucPublicKey,&pucKey,&phKeyHandle)
	key = ConvertToECCCipherGo(pucKey)
	keyHandle = KeyHandleType(phKeyHandle)
	err = ToError(err1)
	return key,keyHandle,err
}
//21.导入会话密钥并用内部 ＥＣＣ私钥解密
func (c *Ctx)SDFImportKeyWithISK_ECC(sessionHandle SessionHandleType,uiISKIndex uint)(core.ECCCipher,KeyHandleType,error){
	var err C.SGD_RV
	var pucKey C.ECCCipher
	var phKeyHandle C.SGD_HANDLE
	err = C.SDFImportKeyWithISK_ECC(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiISKIndex),&pucKey,&phKeyHandle)
	key := ConvertToECCCipherGo(pucKey)
	return key,KeyHandleType(phKeyHandle),ToError(err)
}
//22.生成密钥协商参数并输出
func (c *Ctx)SDFGenerateAgreementDataWithECC(sessionHandle SessionHandleType,uiISKIndex uint,uiKeyBits uint,uiSponsorIDLength uint)([]byte,core.ECCrefPublicKey,core.ECCrefPublicKey,AgreementHandleType,error){
	var err C.SGD_RV
	var pucSponsorID C.SGD_UCHAR_PRT
	var pucSponsorPublicKey C.ECCrefPublicKey
	var pucSponsorTmpPublicKey C.ECCrefPublicKey
	var phAgreementHandle C.SGD_HANDLE
	err = C.SDFGenerateAgreementDataWithECC(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiISKIndex),C.SGD_UINT32(uiKeyBits),&pucSponsorID,C.SGD_UINT32(uiSponsorIDLength),&pucSponsorPublicKey,&pucSponsorTmpPublicKey,&phAgreementHandle)
	sponsorPublicKey :=ConvertToECCrefPublicKeyGo(pucSponsorPublicKey)
	sponsorTmpPublicKey :=ConvertToECCrefPublicKeyGo(pucSponsorTmpPublicKey)
	sponsorID := C.GoBytes(unsafe.Pointer(pucSponsorID), C.int(uiSponsorIDLength))
	C.free(unsafe.Pointer(pucSponsorID))
	return sponsorID,sponsorPublicKey,sponsorTmpPublicKey,AgreementHandleType(phAgreementHandle),ToError(err)
}
//23.计算会话密钥
func (c *Ctx)SDFGenerateKeyWithECC(sessionHandle SessionHandleType,hAgreementHandle AgreementHandleType)([]byte,core.ECCrefPublicKey,core.ECCrefPublicKey,KeyHandleType,error){
	var err C.SGD_RV
	var pucResponseID C.SGD_UCHAR_PRT
	var uiResponseIDLength C.SGD_UINT32
	var pucResponsePublicKey C.ECCrefPublicKey
	var pucResponseTmpPublicKey C.ECCrefPublicKey
	var phKeyHandle C.SGD_HANDLE
	err = C.SDFGenerateKeyWithECC(c.libHandle,C.SGD_HANDLE(sessionHandle),&pucResponseID,uiResponseIDLength,&pucResponsePublicKey,&pucResponseTmpPublicKey,C.SGD_HANDLE(hAgreementHandle),&phKeyHandle)
	responsePublicKey :=ConvertToECCrefPublicKeyGo(pucResponsePublicKey)
	responseTmpPublicKey :=ConvertToECCrefPublicKeyGo(pucResponseTmpPublicKey)
	responseID := C.GoBytes(unsafe.Pointer(pucResponseID), C.int(uiResponseIDLength))
	C.free(unsafe.Pointer(pucResponseID))
	return responseID,responsePublicKey,responseTmpPublicKey,KeyHandleType(phKeyHandle),ToError(err)
}
//24.产生协商数据并计算会话密钥
func (c *Ctx)SDFGenerateAgreementDataAndKeyWithECC(sessionHandle SessionHandleType,uiISKIndex uint,uiKeyBits uint,uiResponseIDLength uint,uiSponsorIDLength uint)([]byte,[]byte,core.ECCrefPublicKey,core.ECCrefPublicKey,core.ECCrefPublicKey,core.ECCrefPublicKey,KeyHandleType,error){
	var err C.SGD_RV
	var pucResponseID C.SGD_UCHAR_PRT
	var pucSponsorID C.SGD_UCHAR_PRT
	var pucSponsorPublicKey C.ECCrefPublicKey
	var pucSponsorTmpPublicKey C.ECCrefPublicKey
	var pucResponsePublicKey C.ECCrefPublicKey
	var pucResponseTmpPublicKey C.ECCrefPublicKey
    var phKeyHandle C.SGD_HANDLE
	err = C.SDFGenerateAgreementDataAndKeyWithECC(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiISKIndex),C.SGD_UINT32(uiKeyBits),&pucResponseID,C.SGD_UINT32(uiResponseIDLength),&pucSponsorID,C.SGD_UINT32(uiSponsorIDLength),&pucSponsorPublicKey,&pucSponsorTmpPublicKey,&pucResponsePublicKey,&pucResponseTmpPublicKey,&phKeyHandle)
	sponsorPublicKey :=ConvertToECCrefPublicKeyGo(pucSponsorPublicKey)
	sponsorTmpPublicKey :=ConvertToECCrefPublicKeyGo(pucSponsorTmpPublicKey)
	responsePublicKey :=ConvertToECCrefPublicKeyGo(pucResponsePublicKey)
	responseTmpPublicKey :=ConvertToECCrefPublicKeyGo(pucResponseTmpPublicKey)
	responseID := C.GoBytes(unsafe.Pointer(pucResponseID), C.int(uiResponseIDLength))
	C.free(unsafe.Pointer(pucResponseID))
	sponsorID := C.GoBytes(unsafe.Pointer(pucSponsorID), C.int(uiSponsorIDLength))
	C.free(unsafe.Pointer(pucSponsorID))
	return responseID,sponsorID,sponsorPublicKey,sponsorTmpPublicKey,responsePublicKey,responseTmpPublicKey,KeyHandleType(phKeyHandle),ToError(err)
}
//25.基于 ＥＣＣ算法的数字信封转换
func (c *Ctx)SDFExchangeDigitEnvelopeBaseOnECC(sessionHandle SessionHandleType,uiKeyIndex uint,uiAlgID uint)(core.ECCrefPublicKey,core.ECCCipher,core.ECCCipher,error){
	var err C.SGD_RV
	var pucPublicKey C.ECCrefPublicKey
	var pucEncDataIn C.ECCCipher
	var pucEncDataOut C.ECCCipher
	err = C.SDFExchangeDigitEnvelopeBaseOnECC(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyIndex),C.SGD_UINT32(uiAlgID),&pucPublicKey,&pucEncDataIn,&pucEncDataOut)
	publicKey :=ConvertToECCrefPublicKeyGo(pucPublicKey)
	encDataIn := ConvertToECCCipherGo(pucEncDataIn)
	encDataOut := ConvertToECCCipherGo(pucEncDataOut)
	return publicKey,encDataIn,encDataOut,ToError(err)
}
//26.生成会话密钥并用密钥加密密钥加密输出
func (c *Ctx)SDFGenerateKeyWithKEK(sessionHandle SessionHandleType,uiKeyBits uint,uiAlgID uint,uiKEKIndex uint )([]byte,uint,KeyHandleType,error){
	var err C.SGD_RV
	var pucKey C.SGD_UCHAR_PRT
	var keyLength C.SGD_UINT32
	var phKeyHandle C.SGD_HANDLE
	err = C.SDFGenerateKeyWithKEK(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyBits),C.SGD_UINT32(uiAlgID),C.SGD_UINT32(uiKEKIndex),&pucKey,&keyLength,&phKeyHandle)
	p:= C.GoBytes(unsafe.Pointer(pucKey), C.int(keyLength))
	C.free(unsafe.Pointer(pucKey))
	return p,uint(keyLength),KeyHandleType(phKeyHandle),ToError(err)
}
//27.导入会话密钥并用密钥加密密钥解密
func (c *Ctx)SDFImportKeyWithKEK(sessionHandle SessionHandleType,uiAlgID uint,uiKEKIndex uint,uiKeyLength uint )([]byte,KeyHandleType,error){
	var err C.SGD_RV
	var pucKey C.SGD_UCHAR_PRT
	var phKeyHandle C.SGD_HANDLE
	err = C.SDFImportKeyWithKEK(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiAlgID),C.SGD_UINT32(uiKEKIndex),&pucKey,C.SGD_UINT32(uiKeyLength),&phKeyHandle)
	p:= C.GoBytes(unsafe.Pointer(pucKey), C.int(uiKeyLength))
	C.free(unsafe.Pointer(pucKey))
	return p,KeyHandleType(phKeyHandle),ToError(err)
}
//28.导入明文会话密钥
func (c *Ctx)SDFImportKey(sessionHandle SessionHandleType,pucKey []byte,uiKeyLength uint)(keyHandle KeyHandleType,err error){
	var err1 C.SGD_RV
	var phKeyHandle C.SGD_HANDLE
	err1 = C.SDFImportKey(c.libHandle,C.SGD_HANDLE(sessionHandle),CMessage(pucKey),C.SGD_UINT32(uiKeyLength),&phKeyHandle)
	keyHandle = KeyHandleType(phKeyHandle)
	err=ToError(err1)
	return keyHandle,err
}
//29.销毁会话密钥
func (c *Ctx)SDFDestroyKey(sessionHandle SessionHandleType,hKeyHandle KeyHandleType)(err error){
	var err1 C.SGD_RV
	err1 = C.SDFDestroyKey(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_HANDLE(hKeyHandle))
	err = ToError(err1)
	return err
}

//30.外部公钥 ＲＳＡ 运算
func (c *Ctx)SDFExternalPublicKeyOperation_RSA(sessionHandle SessionHandleType,uiInputLength uint)(core.RSArefPublicKey,[]byte,[]byte,error){
	var err C.SGD_RV
	var pucPublicKey C.RSArefPublicKey
	var pucDataInput C.SGD_UCHAR_PRT
	var pucDataOutput C.SGD_UCHAR_PRT
	var puiOutputLength C.SGD_UINT32
	err = C.SDFExternalPublicKeyOperation_RSA(c.libHandle,C.SGD_HANDLE(sessionHandle),&pucPublicKey,&pucDataInput,C.SGD_UINT32(uiInputLength),&pucDataOutput,&puiOutputLength)
	publicKey := ConvertToRSArefPublicKeyGo(pucPublicKey)
	dataInput:= C.GoBytes(unsafe.Pointer(pucDataInput), C.int(uiInputLength))
	C.free(unsafe.Pointer(pucDataInput))
	dataOutput:= C.GoBytes(unsafe.Pointer(pucDataOutput), C.int(puiOutputLength))
	C.free(unsafe.Pointer(pucDataOutput))
	return publicKey,dataInput,dataOutput,ToError(err)
}

//31. 外部私钥ＲＳＡ运算
func (c *Ctx)SDFExternalPrivateKeyOperation_RSA(sessionHandle SessionHandleType,uiInputLength uint)(core.RSArefPrivateKey,[]byte,[]byte,error){
	var err C.SGD_RV
	var pucPrivateKey C.RSArefPrivateKey
	var pucDataInput C.SGD_UCHAR_PRT
	var pucDataOutput C.SGD_UCHAR_PRT
	var puiOutputLength C.SGD_UINT32
	err = C.SDFExternalPrivateKeyOperation_RSA(c.libHandle,C.SGD_HANDLE(sessionHandle),&pucPrivateKey,&pucDataInput,C.SGD_UINT32(uiInputLength),&pucDataOutput,&puiOutputLength)
	privateKey :=ConvertToRSArefPrivateKeyGo(pucPrivateKey)
	dataInput:= C.GoBytes(unsafe.Pointer(pucDataInput), C.int(uiInputLength))
	C.free(unsafe.Pointer(pucDataInput))
	dataOutput:= C.GoBytes(unsafe.Pointer(pucDataOutput), C.int(puiOutputLength))
	C.free(unsafe.Pointer(pucDataOutput))
	return privateKey,dataInput,dataOutput,ToError(err)
}

//32.内部公钥 ＲＳＡ 运算
func (c *Ctx)SDFInternalPublicKeyOperation_RSA(sessionHandle SessionHandleType,uiKeyIndex uint,pucDataInput []byte,uiInputLength uint)(dataOutput []byte,err error){
	var err1 C.SGD_RV
	var pucDataOutput C.SGD_UCHAR_PRT
	var puiOutputLength C.SGD_UINT32
	err1 = C.SDFInternalPublicKeyOperation_RSA(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyIndex),CMessage(pucDataInput),C.SGD_UINT32(uiInputLength),&pucDataOutput,&puiOutputLength)
	dataOutput = C.GoBytes(unsafe.Pointer(pucDataOutput), C.int(puiOutputLength))
	err=ToError(err1)
	//C.free(unsafe.Pointer(pucDataOutput))
	return dataOutput,err
}

//33.外部私钥 ＲＳＡ 运算
func (c *Ctx)SDFInternalPrivateKeyOperation_RSA(sessionHandle SessionHandleType,uiKeyIndex uint,inData []byte,uiInputLength uint)(dataOutput []byte,err error){
	var err1 C.SGD_RV
	var pucDataOutput C.SGD_UCHAR_PRT
	var puiOutputLength C.SGD_UINT32
	err1 = C.SDFInternalPrivateKeyOperation_RSA(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyIndex),CMessage(inData),C.SGD_UINT32(uiInputLength),&pucDataOutput,&puiOutputLength)
	dataOutput1 := C.GoBytes(unsafe.Pointer(pucDataOutput), C.int(puiOutputLength))
	dataOutput = deepCopy(dataOutput1)
	//C.free(unsafe.Pointer(pucDataOutput))
	err = ToError(err1)
	return dataOutput,err
}

//34. 外部密钥ＥＣＣ签名
func (c *Ctx)SDFExternalSign_ECC(sessionHandle SessionHandleType,uiAlgID uint,privateKey core.ECCrefPrivateKey,pucData []byte,uiDataLength uint)(signature core.ECCSignature,err error) {
	var err1 C.SGD_RV
	var pucPrivateKey C.ECCrefPrivateKey
	pucPrivateKey.bits = C.SGD_UINT32(privateKey.Bits)
	for i:=0;i<len(privateKey.K);i++{
		pucPrivateKey.K[i]=C.SGD_UCHAR(privateKey.K[i])
	}
	var pucSignature C.ECCSignature
	err1 = C.SDFExternalSign_ECC(c.libHandle, C.SGD_HANDLE(sessionHandle), C.SGD_UINT32(uiAlgID), &pucPrivateKey, CMessage(pucData), C.SGD_UINT32(uiDataLength), &pucSignature)
	signature = ConvertToECCSignatureGo(pucSignature)
	err = ToError(err1)
	return signature,err
}

//35.外部密钥 ＥＣＣ验证
func (c *Ctx)SDFExternalVerify_ECC(sessionHandle SessionHandleType,uiAlgID uint,uiInputLength uint)(core.ECCrefPublicKey,core.ECCSignature,[]byte,error){
	var err C.SGD_RV
	var pucPublicKey C.ECCrefPublicKey
	var pucSignature C.ECCSignature
	var pucDataInput C.SGD_UCHAR_PRT
	err = C.SDFExternalVerify_ECC(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiAlgID),&pucPublicKey,&pucDataInput,C.SGD_UINT32(uiInputLength),&pucSignature)
	publicKey := ConvertToECCrefPublicKeyGo(pucPublicKey)
	signature := ConvertToECCSignatureGo(pucSignature)
	input:= C.GoBytes(unsafe.Pointer(pucDataInput), C.int(uiInputLength))
	C.free(unsafe.Pointer(pucDataInput))
	return publicKey,signature,input,ToError(err)
}
//36.内部密钥 ＥＣＣ签名
func (c *Ctx)SDFInternalSign_ECC(sessionHandle SessionHandleType,uiISKIndex uint,pucData []byte,uiDataLength uint)(signature core.ECCSignature,err error){
	var err1 C.SGD_RV
	var pucSignature C.ECCSignature
	err1 = C.SDFInternalSign_ECC(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiISKIndex),CMessage(pucData),C.SGD_UINT32(uiDataLength),&pucSignature)
	signature = ConvertToECCSignatureGo(pucSignature)
	err = ToError(err1)
	return signature,err
}
//37.内部密钥 ＥＣＣ验证
func (c *Ctx)SDFInternalVerify_ECC(sessionHandle SessionHandleType,uiISKIndex uint,pucData []byte,uiDataLength uint,signature core.ECCSignature)(err error){
	var err1 C.SGD_RV
	var pucSignature C.ECCSignature
	for i:=0;i<len(signature.R);i++{
		pucSignature.r[i]=C.SGD_UCHAR(signature.R[i])
	}
	for i:=0;i<len(signature.S);i++{
		pucSignature.s[i]=C.SGD_UCHAR(signature.S[i])
	}
	err1 = C.SDFInternalVerify_ECC(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiISKIndex),CMessage(pucData),C.SGD_UINT32(uiDataLength),&pucSignature)
	err = ToError(err1)
	return err
}
//38.外部密钥 ＥＣＣ加密
func (c *Ctx)SDFExternalEncrypt_ECC(sessionHandle SessionHandleType,uiAlgID uint,publicKey core.ECCrefPublicKey,data []byte,dataLength uint)(encData core.ECCCipher,err error){
	var err1 C.SGD_RV
	var pucPublicKey C.ECCrefPublicKey
	pucPublicKey.bits=C.SGD_UINT32(publicKey.Bits)
	for i:=0;i<len(publicKey.X);i++{
		pucPublicKey.x[i]=C.SGD_UCHAR(publicKey.X[i])
	}
	for i:=0;i<len(publicKey.Y);i++{
		pucPublicKey.y[i]=C.SGD_UCHAR(publicKey.Y[i])
	}
	var pucEncData C.ECCCipher
	err1 = C.SDFExternalEncrypt_ECC(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiAlgID),&pucPublicKey,CMessage(data),C.SGD_UINT32(dataLength),&pucEncData)
	encData = ConvertToECCCipherGo(pucEncData)
	err = ToError(err1)
	return encData,err
}
//39.外部密钥 ＥＣＣ解密
func (c *Ctx)SDFExternalDecrypt_ECC(sessionHandle SessionHandleType,uiAlgID uint,privateKey core.ECCrefPrivateKey,encData core.ECCCipher)(data []byte,dataLength uint,err error){
	var err1 C.SGD_RV
	var pucPrivateKey C.ECCrefPrivateKey
	pucPrivateKey.bits=C.SGD_UINT32(privateKey.Bits)
	for i:=0;i<len(privateKey.K);i++{
		pucPrivateKey.K[i]=C.SGD_UCHAR(privateKey.K[i])
	}
	var pucEncData C.ECCCipher
	for i:=0;i<len(encData.X);i++{
		pucEncData.x[i]=C.SGD_UCHAR(encData.X[i])
	}
	for i:=0;i<len(encData.Y);i++{
		pucEncData.y[i]=C.SGD_UCHAR(encData.Y[i])
	}
	for i:=0;i<len(encData.M);i++{
		pucEncData.M[i]=C.SGD_UCHAR(encData.M[i])
	}
	pucEncData.L=C.SGD_UINT32(encData.L)
	for i:=0;i<len(encData.C);i++{
		pucEncData.C[i]=C.SGD_UCHAR(encData.C[i])
	}
	var pucData C.SGD_UCHAR_PRT
	var puiDataLength C.SGD_UINT32
	err1 = C.SDFExternalDecrypt_ECC(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiAlgID), &pucPrivateKey,&pucEncData,&pucData,&puiDataLength)
	data = C.GoBytes(unsafe.Pointer(pucData), C.int(puiDataLength))
	dataLength=uint(puiDataLength)
	//C.free(unsafe.Pointer(pucData))
	err = ToError(err1)
	return data,dataLength,err
}

//40.对称加密
func (c *Ctx)SDFEncrypt(sessionHandle SessionHandleType,keyHandle KeyHandleType,algID uint,iv []byte,data []byte, dataLength uint)(encData []byte,encDataLength uint,err error){
	var err1 C.SGD_RV
	var pucEncData C.SGD_UCHAR_PRT
	var puiEncDataLength C.SGD_UINT32
	err1 = C.SDFEncrypt(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_HANDLE(keyHandle),C.SGD_UINT32(algID),CMessage(iv),CMessage(data),C.SGD_UINT32(dataLength),&pucEncData,&puiEncDataLength)
	encData = C.GoBytes(unsafe.Pointer(pucEncData), C.int(puiEncDataLength))
	encDataLength = uint(puiEncDataLength)
	err = ToError(err1)
	//C.free(unsafe.Pointer(pucEncData))
	return encData,uint(puiEncDataLength),err
}
//41.对称解密
func (c *Ctx)SDFDecrypt(sessionHandle SessionHandleType,hKeyHandle KeyHandleType,uiAlgID uint,iv []byte, encData []byte,encDataLength uint)(data []byte,dataLength uint ,err error){
	var err1 C.SGD_RV
	var pucData C.SGD_UCHAR_PRT
	var puiDataLength C.SGD_UINT32
	err1 = C.SDFDecrypt(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_HANDLE(hKeyHandle),C.SGD_UINT32(uiAlgID),CMessage(iv),CMessage(encData),C.SGD_UINT32(encDataLength),&pucData,&puiDataLength)
	data = C.GoBytes(unsafe.Pointer(pucData), C.int(puiDataLength))
	dataLength = uint(puiDataLength)
	//C.free(unsafe.Pointer(pucData))
	err = ToError(err1)
	return data,dataLength,err
}
//42.计算 ＭＡＣ
func (c *Ctx)SDFCalculateMAC(sessionHandle SessionHandleType,hKeyHandle KeyHandleType,uiAlgID uint,iv []byte, data []byte,dataLength uint)(mac []byte,macLength uint,err error){
	var err1 C.SGD_RV
	var pucMAC C.SGD_UCHAR_PRT
	var puiMACLength C.SGD_UINT32
	err1 = C.SDFCalculateMAC(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_HANDLE(hKeyHandle),C.SGD_UINT32(uiAlgID),CMessage(iv),CMessage(data),C.SGD_UINT32(dataLength),&pucMAC,&puiMACLength)
	mac= C.GoBytes(unsafe.Pointer(pucMAC), C.int(puiMACLength))
	macLength = uint(puiMACLength)
	//C.free(unsafe.Pointer(pucMAC))
	err = ToError(err1)
	return mac,macLength,err
}
//43.杂凑运算初始化
func (c *Ctx)SDFHashInit(sessionHandle SessionHandleType,uiAlgID uint,pucID []byte,uiIDLength uint)(publicKey core.ECCrefPublicKey,err error){
	var err1 C.SGD_RV
	var pucPublicKey C.ECCrefPublicKey
	err1 = C.SDFHashInit(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiAlgID),&pucPublicKey,CMessage(pucID),C.SGD_UINT32(uiIDLength))
	publicKey = ConvertToECCrefPublicKeyGo(pucPublicKey)
	err =ToError(err1)
	return publicKey,err
}
//44.多包杂凑运算
func (c *Ctx)SDFHashUpdate(sessionHandle SessionHandleType,pucData []byte,uiDataLength uint)(err error){
	var err1 C.SGD_RV
	err1 = C.SDFHashUpdate(c.libHandle,C.SGD_HANDLE(sessionHandle),CMessage(pucData),C.SGD_UINT32(uiDataLength))
	err =ToError(err1)
	return err
}
//45.杂凑运算结束
func (c *Ctx)SDFHashFinal(sessionHandle SessionHandleType)(hash []byte,hashLength uint,err error){
	var err1 C.SGD_RV
	var pucData C.SGD_UCHAR_PRT
	var puiHashLength C.SGD_UINT32
	err1 = C.SDFHashFinal(c.libHandle,C.SGD_HANDLE(sessionHandle),&pucData,&puiHashLength)
	hash = C.GoBytes(unsafe.Pointer(pucData), C.int(puiHashLength))
	hashLength = uint(puiHashLength)
	//C.free(unsafe.Pointer(pucData))
	err = ToError(err1)
	return hash,hashLength,err
}
//46.创建文件
func (c *Ctx)SDFCreateFile(sessionHandle SessionHandleType,fileName []byte,uiFileSize uint)(err error){
	var err1 C.SGD_RV
	err1 = C.SDFCreateFile(c.libHandle,C.SGD_HANDLE(sessionHandle),CMessage(fileName),C.SGD_UINT32(len(fileName)),C.SGD_UINT32(uiFileSize))
	err = ToError(err1)
	return err
}
//47.读取文件
func (c *Ctx)SDFReadFile(sessionHandle SessionHandleType,fileName []byte,uiOffset uint)(buffer []byte,readLength uint,err error){
	var err1 C.SGD_RV
	var puiReadLength C.SGD_UINT32
	var pucBuffer C.SGD_UCHAR_PRT
	err1 = C.SDFReadFile(c.libHandle,C.SGD_HANDLE(sessionHandle),CMessage(fileName),C.SGD_UINT32(len(fileName)),C.SGD_UINT32(uiOffset),&puiReadLength,&pucBuffer)
	buffer = C.GoBytes(unsafe.Pointer(pucBuffer), C.int(puiReadLength))
	readLength = uint(puiReadLength)
	//C.free(unsafe.Pointer(pucBuffer))
	err = ToError(err1)
	return buffer,readLength,err
}
//48.写文件
func (c *Ctx)SDFWriteFile(sessionHandle SessionHandleType,fileName []byte,uiOffset uint ,pucBuffer []byte)(err error){
	var err1 C.SGD_RV
	err1 = C.SDFWriteFile(c.libHandle,C.SGD_HANDLE(sessionHandle),CMessage(fileName),C.SGD_UINT32(len(fileName)),C.SGD_UINT32(uiOffset),C.SGD_UINT32(len(pucBuffer)),CMessage(pucBuffer))
	err = ToError(err1)
	return err
}
//49.删除文件
func (c *Ctx)SDFDeleteFile(sessionHandle SessionHandleType,fileName []byte)(err error){
	var err1 C.SGD_RV
	err1 = C.SDFDeleteFile(c.libHandle,C.SGD_HANDLE(sessionHandle),CMessage(fileName),C.SGD_UINT32(len(fileName)))
	err = ToError(err1)
	return err
}
//50.
func (c *Ctx)SDFGetSymmKeyHandle(sessionHandle SessionHandleType,uiKeyIndex uint)(KeyHandleType,error){
	var err C.SGD_RV
	var phKeyHandle C.SGD_HANDLE
	err = C.SDFGetSymmKeyHandle(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyIndex),&phKeyHandle)
	return KeyHandleType(phKeyHandle),ToError(err)
}
//51. ECC方式的加密
func (c *Ctx)SDFInternalEncrypt_ECC(sessionHandle SessionHandleType,uiISKIndex uint,uiAlgID uint,pucData []byte,uiDataLength uint)(encData core.ECCCipher,err error){
	var err1 C.SGD_RV
	var pucEncData C.ECCCipher
	err1 = C.SDFInternalEncrypt_ECC(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiISKIndex),C.SGD_UINT32(uiAlgID),CMessage(pucData),C.SGD_UINT32(uiDataLength),&pucEncData)
	encData =ConvertToECCCipherGo(pucEncData)
	err = ToError(err1)
	return encData,err
}
//52. ECC方式的解密
func (c *Ctx)SDFInternalDecrypt_ECC(sessionHandle SessionHandleType,uiISKIndex uint,uiAlgID uint,encData core.ECCCipher)(data []byte,dataLength uint,err error){
	var err1 C.SGD_RV
	var pucEncData C.ECCCipher
	for i:=0;i<len(encData.X);i++{
		pucEncData.x[i]=C.SGD_UCHAR(encData.X[i])
	}
	for i:=0;i<len(encData.Y);i++{
		pucEncData.y[i]=C.SGD_UCHAR(encData.Y[i])
	}
	for i:=0;i<len(encData.M);i++{
		pucEncData.M[i]=C.SGD_UCHAR(encData.M[i])
	}
	pucEncData.L=C.SGD_UINT32(encData.L)
	for i:=0;i<len(encData.C);i++{
		pucEncData.C[i]=C.SGD_UCHAR(encData.C[i])
	}
	var pucData C.SGD_UCHAR_PRT
	var puiDataLength C.SGD_UINT32
	err1 = C.SDFInternalDecrypt_ECC(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiISKIndex),C.SGD_UINT32(uiAlgID),&pucEncData,&pucData,&puiDataLength)
	//encData :=ConvertToECCCipherGo(pucEncData)
	data= C.GoBytes(unsafe.Pointer(pucData), C.int(puiDataLength))
	//C.free(unsafe.Pointer(pucData))
	dataLength = uint(puiDataLength)
	err = ToError(err1)
	return data,dataLength,err
}

//53. EPK方式导出RSA密钥
func (c *Ctx)SDFExportKeyWithEPK_RSA(sessionHandle SessionHandleType,hKeyHandle KeyHandleType)(core.RSArefPublicKey,[]byte,error){
	var err C.SGD_RV
	var pucPublicKey C.RSArefPublicKey
	var pucKey C.SGD_UCHAR_PRT
	var puiKeyLength C.SGD_UINT32
	err = C.SDFExportKeyWithEPK_RSA(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_HANDLE(hKeyHandle),&pucPublicKey,&pucKey,&puiKeyLength)
	publicKey := ConvertToRSArefPublicKeyGo(pucPublicKey)
	key:= C.GoBytes(unsafe.Pointer(pucKey), C.int(puiKeyLength))
	C.free(unsafe.Pointer(pucKey))
	return publicKey,key,ToError(err)
}

//54. EPK方式导出ECC密钥
func (c *Ctx)SDFExportKeyWithEPK_ECC(sessionHandle SessionHandleType,hKeyHandle KeyHandleType,uiAlgID uint)(core.ECCrefPublicKey,core.ECCCipher,error){
	var err C.SGD_RV
	var pucPublicKey C.ECCrefPublicKey
	var pucKey C.ECCCipher
	err = C.SDFExportKeyWithEPK_ECC(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_HANDLE(hKeyHandle),C.SGD_UINT32(uiAlgID),&pucPublicKey,&pucKey)
	publicKey := ConvertToECCrefPublicKeyGo(pucPublicKey)
	key :=ConvertToECCCipherGo(pucKey)

	return publicKey,key,ToError(err)
}
//55. EPK方式导出密钥
func (c *Ctx)SDFExportKeyWithKEK(sessionHandle SessionHandleType,hKeyHandle KeyHandleType,uiAlgID uint,uiKEKIndex uint)([]byte,error){
	var err C.SGD_RV
	var pucKey C.SGD_UCHAR_PRT
	var puiKeyLength C.SGD_UINT32
	err = C.SDFExportKeyWithKEK(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_HANDLE(hKeyHandle),C.SGD_UINT32(uiAlgID),C.SGD_UINT32(uiKEKIndex),&pucKey,&puiKeyLength)
	key:= C.GoBytes(unsafe.Pointer(pucKey), C.int(puiKeyLength))
	C.free(unsafe.Pointer(pucKey))
	return key,ToError(err)
}
//56. 导出SM9签名主公钥
func (c *Ctx)SDFExportSignMasterPublicKey_SM9(sessionHandle SessionHandleType,uiKeyIndex uint)(core.SM9refSignMasterPublicKey,error){
	var err C.SGD_RV
	var pPublicKey C.SM9refSignMasterPublicKey
	err = C.SDFExportSignMasterPublicKey_SM9(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyIndex),&pPublicKey)
	publicKey := ConvertToSM9refSignMasterPublicKeyGo(pPublicKey)
	return publicKey,ToError(err)
}
//57. 导出SM9加密主公钥
func (c *Ctx)SDFExportEncMasterPublicKey_SM9(sessionHandle SessionHandleType,uiKeyIndex uint)(core.SM9refEncMasterPublicKey,error){
	var err C.SGD_RV
	var pEncMasterPublicKey C.SM9refEncMasterPublicKey
	err = C.SDFExportEncMasterPublicKey_SM9(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyIndex),&pEncMasterPublicKey)
	publicKey := ConvertToSM9refEncMasterPublicKeyGo(pEncMasterPublicKey)
	return publicKey,ToError(err)
}

//58. 导出SM9签名主密钥对
func (c *Ctx)SDFExportSignMasterKeyPairG_SM9(sessionHandle SessionHandleType,uiKeyIndex uint)([]byte,error){
	var err C.SGD_RV
	var pPairG C.SGD_UCHAR_PRT
	var puiPairGLen C.SGD_UINT32
	err = C.SDFExportSignMasterKeyPairG_SM9(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyIndex),&pPairG,&puiPairGLen)
	pairG:= C.GoBytes(unsafe.Pointer(pPairG), C.int(puiPairGLen))
	C.free(unsafe.Pointer(pPairG))
	return pairG,ToError(err)
}

//59. 导出SM9加密主密钥对
func (c *Ctx)SDFExportEncMasterKeyPairG_SM9(sessionHandle SessionHandleType,uiKeyIndex uint)([]byte,error){
	var err C.SGD_RV
	var pPairG C.SGD_UCHAR_PRT
	var puiPairGLen C.SGD_UINT32
	err = C.SDFExportEncMasterKeyPairG_SM9(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyIndex),&pPairG,&puiPairGLen)
	pairG:= C.GoBytes(unsafe.Pointer(pPairG), C.int(puiPairGLen))
	C.free(unsafe.Pointer(pPairG))
	return pairG,ToError(err)
}
//60. 导入SM9使用者签名私钥
func (c *Ctx)SDFImportUserSignPrivateKey_SM9(sessionHandle SessionHandleType,uiKeyIndex uint)(core.SM9refSignUserPrivateKey,error){
	var err C.SGD_RV
	var pUserPrivateKey C.SM9refSignUserPrivateKey
	err = C.SDFImportUserSignPrivateKey_SM9(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyIndex),&pUserPrivateKey)
	privateKey:=ConvertToSM9refSignUserPrivateKeyGo(pUserPrivateKey)
	return privateKey,ToError(err)
}

//61. 导入SM9使用者加密私钥
func (c *Ctx)SDFImportUserEncPrivateKey_SM9(sessionHandle SessionHandleType,uiKeyIndex uint)(core.SM9refEncUserPrivateKey,error){
	var err C.SGD_RV
	var pUserPrivateKey C.SM9refEncUserPrivateKey
	err = C.SDFImportUserEncPrivateKey_SM9(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyIndex),&pUserPrivateKey)
	privateKey := ConvertToSM9refEncUserPrivateKeyGo(pUserPrivateKey)
	return privateKey,ToError(err)
}
//62. 产生SM9使用者签名私钥
func (c *Ctx)SDFGenerateSignUserPrivateKey_SM9(sessionHandle SessionHandleType,uiKeyIndex uint)(core.SM9refSignUserPrivateKey,error){
	var err C.SGD_RV
	var hid C.SGD_UCHAR
	var pucUserID C.SGD_UCHAR_PRT
	var uiUserIDLen C.SGD_UINT32
	var pUserPrivateKey C.SM9refSignUserPrivateKey
	err = C.SDFGenerateSignUserPrivateKey_SM9(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyIndex),hid,&pucUserID,C.SGD_UINT32(uiUserIDLen),&pUserPrivateKey)
	privateKey:=ConvertToSM9refSignUserPrivateKeyGo(pUserPrivateKey)
	return privateKey,ToError(err)
}

//63. 产生SM9使用者加密私钥
func (c *Ctx)SDFGenerateEncUserPrivateKey_SM9(sessionHandle SessionHandleType,uiKeyIndex uint,uiUserIDLen uint)(core.SM9refEncUserPrivateKey,[]byte,error){
	var err C.SGD_RV
	var hid C.SGD_UCHAR
	var pucUserID C.SGD_UCHAR_PRT
	var pUserPrivateKey C.SM9refEncUserPrivateKey
	err = C.SDFGenerateEncUserPrivateKey_SM9(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyIndex),hid,&pucUserID,C.SGD_UINT32(uiUserIDLen),&pUserPrivateKey)
	privateKey := ConvertToSM9refEncUserPrivateKeyGo(pUserPrivateKey)
	userID:= C.GoBytes(unsafe.Pointer(pucUserID), C.int(uiUserIDLen))
	C.free(unsafe.Pointer(pucUserID))
	return privateKey,userID,ToError(err)
}

//64.SM9签名
func (c *Ctx)SDFSign_SM9(sessionHandle SessionHandleType,uiKeyIndex uint)(core.SM9refSignUserPrivateKey,core.SM9refSignMasterPublicKey,[]byte,core.SM9Signature,error){
	var err C.SGD_RV
	var pUserPrivateKey C.SM9refSignUserPrivateKey
	var pMasterPublicKey C.SM9refSignMasterPublicKey
	var pucDataInput C.SGD_UCHAR_PRT
	var uiDataInputLen C.SGD_UINT32
	var pSignature C.SM9Signature
	err = C.SDFSign_SM9(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyIndex),&pUserPrivateKey,&pMasterPublicKey,&pucDataInput,uiDataInputLen,&pSignature)
	privateKey:=ConvertToSM9refSignUserPrivateKeyGo(pUserPrivateKey)
	publicKey := ConvertToSM9refSignMasterPublicKeyGo(pMasterPublicKey)
	sign := ConvertToSM9SignatureGo(pSignature)
	dataInput:= C.GoBytes(unsafe.Pointer(pucDataInput), C.int(uiDataInputLen))
	C.free(unsafe.Pointer(pucDataInput))
	return privateKey,publicKey,dataInput,sign,ToError(err)
}
//65. SM9签名扩展方法
func (c *Ctx)SDFSignEx_SM9(sessionHandle SessionHandleType,uiKeyIndex uint,uiPairGLen uint,uiDataInputLen uint)(core.SM9refSignUserPrivateKey,core.SM9refSignMasterPublicKey,[]byte,[]byte,core.SM9Signature,error){
	var err C.SGD_RV
	var pUserPrivateKey C.SM9refSignUserPrivateKey
	var pMasterPublicKey C.SM9refSignMasterPublicKey
	var pPairG C.SGD_UCHAR_PRT
	var pucDataInput C.SGD_UCHAR_PRT
	var pSignature C.SM9Signature
	err = C.SDFSignEx_SM9(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyIndex),&pUserPrivateKey,&pMasterPublicKey,&pPairG,C.SGD_UINT32(uiPairGLen),&pucDataInput,C.SGD_UINT32(uiDataInputLen),&pSignature)
	privateKey:=ConvertToSM9refSignUserPrivateKeyGo(pUserPrivateKey)
	publicKey := ConvertToSM9refSignMasterPublicKeyGo(pMasterPublicKey)
	sign := ConvertToSM9SignatureGo(pSignature)
	pairGLen:= C.GoBytes(unsafe.Pointer(pPairG), C.int(uiPairGLen))
	C.free(unsafe.Pointer(pPairG))
	dataInput:= C.GoBytes(unsafe.Pointer(pucDataInput), C.int(uiDataInputLen))
	C.free(unsafe.Pointer(pucDataInput))
	return privateKey,publicKey,dataInput,pairGLen,sign,ToError(err)
}

//66. SM9验证
func (c *Ctx)SDFVerify_SM9(sessionHandle SessionHandleType,uiUserIDLen uint,uiDataInputLen uint)(core.SM9refSignMasterPublicKey,[]byte,[]byte,core.SM9Signature,error){
	var err C.SGD_RV
	var hid C.SGD_UCHAR
	var pucUserID C.SGD_UCHAR_PRT
	var pMasterPublicKey C.SM9refSignMasterPublicKey
	var pucData C.SGD_UCHAR_PRT
	var pSignature C.SM9Signature
	err = C.SDFVerify_SM9(c.libHandle,C.SGD_HANDLE(sessionHandle),hid,&pucUserID,C.SGD_UINT32(uiUserIDLen),&pMasterPublicKey,&pucData,C.SGD_UINT32(uiDataInputLen),&pSignature)
	publicKey := ConvertToSM9refSignMasterPublicKeyGo(pMasterPublicKey)
	sign := ConvertToSM9SignatureGo(pSignature)
	userIDLen:= C.GoBytes(unsafe.Pointer(pucUserID), C.int(uiUserIDLen))
	C.free(unsafe.Pointer(pucUserID))
	data := C.GoBytes(unsafe.Pointer(pucData), C.int(uiDataInputLen))
	C.free(unsafe.Pointer(pucUserID))
	return publicKey,userIDLen,data,sign,ToError(err)
}
//67. SM9验证扩展方法
func (c *Ctx)SDFVerifyEx_SM9(sessionHandle SessionHandleType,uiUserIDLen uint,uiPairGLen uint,uiDataInputLen uint)(core.SM9refSignMasterPublicKey,[]byte,[]byte,[]byte,core.SM9Signature,error){
	var err C.SGD_RV
	var hid C.SGD_UCHAR
	var pucUserID C.SGD_UCHAR_PRT
	var pMasterPublicKey C.SM9refSignMasterPublicKey
	var pPairG C.SGD_UCHAR_PRT
	var pucData C.SGD_UCHAR_PRT
	var pSignature C.SM9Signature
	err = C.SDFVerifyEx_SM9(c.libHandle,C.SGD_HANDLE(sessionHandle),hid,&pucUserID,C.SGD_UINT32(uiUserIDLen),&pMasterPublicKey,&pPairG,C.SGD_UINT32(uiPairGLen),&pucData,C.SGD_UINT32(uiDataInputLen),&pSignature)
	publicKey := ConvertToSM9refSignMasterPublicKeyGo(pMasterPublicKey)
	sign := ConvertToSM9SignatureGo(pSignature)
	userIDLen:= C.GoBytes(unsafe.Pointer(pucUserID), C.int(uiUserIDLen))
	C.free(unsafe.Pointer(pucUserID))
	pairG:= C.GoBytes(unsafe.Pointer(pPairG), C.int(uiPairGLen))
	C.free(unsafe.Pointer(pPairG))
	data := C.GoBytes(unsafe.Pointer(pucData), C.int(uiDataInputLen))
	C.free(unsafe.Pointer(pucUserID))
	return publicKey,userIDLen,pairG,data,sign,ToError(err)
}

//68. SM9加密
func (c *Ctx)SDFEncrypt_SM9(sessionHandle SessionHandleType,uiUserIDLen uint,uiDataInputLen uint,uiPairGLen uint)(core.SM9refEncMasterPublicKey,[]byte,[]byte,core.SM9Cipher,error){
	var err C.SGD_RV
	var hid C.SGD_UCHAR
	var pucUserID C.SGD_UCHAR_PRT
	var pEncMasterPublicKey C.SM9refEncMasterPublicKey
	var pucData C.SGD_UCHAR_PRT
	var pCipher C.SM9Cipher
	err = C.SDFEncrypt_SM9(c.libHandle,C.SGD_HANDLE(sessionHandle),hid,&pucUserID,C.SGD_UINT32(uiUserIDLen),&pEncMasterPublicKey,&pucData,C.SGD_UINT32(uiDataInputLen),&pCipher)
	publicKey := ConvertToSM9refEncMasterPublicKeyGo(pEncMasterPublicKey)
	cipher := ConvertToSM9CipherGo(pCipher)
	userIDLen:= C.GoBytes(unsafe.Pointer(pucUserID), C.int(uiUserIDLen))
	C.free(unsafe.Pointer(pucUserID))
	data := C.GoBytes(unsafe.Pointer(pucData), C.int(uiDataInputLen))
	C.free(unsafe.Pointer(pucUserID))
	return publicKey,userIDLen,data,cipher,ToError(err)
}

//69. SM9加密扩展方法
func (c *Ctx)SDFEncryptEx_SM9(sessionHandle SessionHandleType,uiUserIDLen uint,uiDataInputLen uint,nPairGLen uint)(core.SM9refEncMasterPublicKey,[]byte,[]byte,[]byte,core.SM9Cipher,error){
	var err C.SGD_RV
	var hid C.SGD_UCHAR
	var pucUserID C.SGD_UCHAR_PRT
	var pEncMasterPublicKey C.SM9refEncMasterPublicKey
	var pPairG C.SGD_UCHAR_PRT
	var pucData C.SGD_UCHAR_PRT
	var pCipher C.SM9Cipher
	err = C.SDFEncryptEx_SM9(c.libHandle,C.SGD_HANDLE(sessionHandle),hid,&pucUserID,C.SGD_UINT32(uiUserIDLen),&pEncMasterPublicKey,&pPairG,C.SGD_UINT32(nPairGLen),&pucData,C.SGD_UINT32(uiDataInputLen),&pCipher)
	publicKey := ConvertToSM9refEncMasterPublicKeyGo(pEncMasterPublicKey)
	cipher := ConvertToSM9CipherGo(pCipher)
	userID:= C.GoBytes(unsafe.Pointer(pucUserID), C.int(uiUserIDLen))
	C.free(unsafe.Pointer(pucUserID))
	pairG := C.GoBytes(unsafe.Pointer(pPairG), C.int(nPairGLen))
	C.free(unsafe.Pointer(pPairG))
	data := C.GoBytes(unsafe.Pointer(pucData), C.int(uiDataInputLen))
	C.free(unsafe.Pointer(pucData))
	return publicKey,userID,pairG,data,cipher,ToError(err)
}
//70. SM9解密
func (c *Ctx)SDFDecrypt_SM9(sessionHandle SessionHandleType,uiUserIDLen uint,uiKeyIndex uint)([]byte,core.SM9refEncUserPrivateKey,core.SM9Cipher,[]byte,error){
	var err C.SGD_RV
	var pucUserID C.SGD_UCHAR_PRT
	var pUserPrivateKey C.SM9refEncUserPrivateKey
	var pCipher C.SM9Cipher
	var pucPlainData C.SGD_UCHAR_PRT
	var uiPlainDataLength C.SGD_UINT32
	err = C.SDFDecrypt_SM9(c.libHandle,C.SGD_HANDLE(sessionHandle),&pucUserID,C.SGD_UINT32(uiUserIDLen),C.SGD_UINT32(uiKeyIndex),&pUserPrivateKey,&pCipher,&pucPlainData,&uiPlainDataLength)
	userPrivateKey := ConvertToSM9refEncUserPrivateKeyGo(pUserPrivateKey)
	cipher := ConvertToSM9CipherGo(pCipher)
	userID:= C.GoBytes(unsafe.Pointer(pucUserID), C.int(uiUserIDLen))
	C.free(unsafe.Pointer(pucUserID))
    plainData:= C.GoBytes(unsafe.Pointer(pucPlainData), C.int(uiPlainDataLength))
	C.free(unsafe.Pointer(pucPlainData))
	return userID,userPrivateKey,cipher,plainData,ToError(err)
}
//71. SM9密钥封装
func (c *Ctx)SDFEncap_SM9(sessionHandle SessionHandleType,uiUserIDLen uint,uiKeyLen uint)([]byte,core.SM9refEncMasterPublicKey,[]byte,core.SM9refKeyPackage,error){
	var err C.SGD_RV
	var hid C.SGD_UCHAR
	var pucUserID C.SGD_UCHAR_PRT
	var pEncMasterPublicKey C.SM9refEncMasterPublicKey
	var pKey C.SGD_UCHAR_PRT
	var pKeyPackage C.SM9refKeyPackage
	err = C.SDFEncap_SM9(c.libHandle,C.SGD_HANDLE(sessionHandle),hid,&pucUserID,C.SGD_UINT32(uiUserIDLen),&pEncMasterPublicKey,C.SGD_UINT32(uiKeyLen),&pKey,&pKeyPackage)
	publicKey := ConvertToSM9refEncMasterPublicKeyGo(pEncMasterPublicKey)
	keyPackage:=ConvertToSM9refKeyPackageGo(pKeyPackage)
	userID:= C.GoBytes(unsafe.Pointer(pucUserID), C.int(uiUserIDLen))
	C.free(unsafe.Pointer(pucUserID))
	key:= C.GoBytes(unsafe.Pointer(pKey), C.int(uiKeyLen))
	C.free(unsafe.Pointer(pKey))
	return userID,publicKey,key,keyPackage,ToError(err)
}
//72. SM9密钥解封
func (c *Ctx)SDFDecap_SM9(sessionHandle SessionHandleType,uiUserIDLen uint,uiKeyIndex uint,uiKeyLen uint)([]byte,core.SM9refEncUserPrivateKey,core.SM9refKeyPackage,[]byte,error){
	var err C.SGD_RV
	var pucUserID C.SGD_UCHAR_PRT
	var pUserPrivateKey C.SM9refEncUserPrivateKey
	var pKeyPackage C.SM9refKeyPackage
	var pucKey C.SGD_UCHAR_PRT
	err = C.SDFDecap_SM9(c.libHandle,C.SGD_HANDLE(sessionHandle),&pucUserID,C.SGD_UINT32(uiUserIDLen),C.SGD_UINT32(uiKeyIndex),&pUserPrivateKey,&pKeyPackage,C.SGD_UINT32(uiKeyLen),&pucKey)
	privateKey := ConvertToSM9refEncUserPrivateKeyGo(pUserPrivateKey)
	keyPackage:=ConvertToSM9refKeyPackageGo(pKeyPackage)
	userID:= C.GoBytes(unsafe.Pointer(pucUserID), C.int(uiUserIDLen))
	C.free(unsafe.Pointer(pucUserID))
	key:= C.GoBytes(unsafe.Pointer(pucKey), C.int(uiKeyLen))
	C.free(unsafe.Pointer(pucKey))
	return userID,privateKey,keyPackage,key,ToError(err)
}


func (c *Ctx)SDFGenerateAgreementDataWithSM9(sessionHandle SessionHandleType,uiResponseIDLength uint)([]byte,core.SM9refEncMasterPublicKey,core.SM9refEncMasterPublicKey,AgreementHandleType,error){
	var err C.SGD_RV
	var hid C.SGD_UCHAR
	var pucResponseID C.SGD_UCHAR_PRT
	var pPublicKey C.SM9refEncMasterPublicKey
	var pucSponsorTmpPublicKey C.SM9refEncMasterPublicKey
	var phAgreementHandle C.SGD_HANDLE
	err = C.SDFGenerateAgreementDataWithSM9(c.libHandle,C.SGD_HANDLE(sessionHandle),hid,&pucResponseID,C.SGD_UINT32(uiResponseIDLength),&pPublicKey,&pucSponsorTmpPublicKey,&phAgreementHandle)
	publicKey := ConvertToSM9refEncMasterPublicKeyGo(pPublicKey)
	sponsorTmpPublicKey := ConvertToSM9refEncMasterPublicKeyGo(pucSponsorTmpPublicKey)
	responseID:= C.GoBytes(unsafe.Pointer(pucResponseID), C.int(uiResponseIDLength))
	C.free(unsafe.Pointer(pucResponseID))
	return responseID,publicKey,sponsorTmpPublicKey,AgreementHandleType(phAgreementHandle),ToError(err)
}


func (c *Ctx)SDFGenerateAgreemetDataAndKeyWithSM9(sessionHandle SessionHandleType,uiKeyLen uint,uiResponseIDLen uint,uiSponsorIDLen uint,uiKeyIndex uint)([]byte,[]byte,core.SM9refEncUserPrivateKey,core.SM9refEncMasterPublicKey,core.SM9refEncMasterPublicKey,core.SM9refEncMasterPublicKey,[]byte,[]byte,KeyHandleType,error){
	var err C.SGD_RV
	var hid C.SGD_UCHAR
	var pucResponseID C.SGD_UCHAR_PRT
	var pucSponsorID C.SGD_UCHAR_PRT
	var pucResponsePrivateKey C.SM9refEncUserPrivateKey
	var pucPublicKey C.SM9refEncMasterPublicKey
	var pucSponsorTmpPublicKey C.SM9refEncMasterPublicKey
	var pucResponseTmpPublicKey C.SM9refEncMasterPublicKey
	var pucHashSB C.SGD_UCHAR_PRT
	var pucHashS2 C.SGD_UCHAR_PRT
	var puiSBLen C.SGD_UINT32
	var puiS2Len C.SGD_UINT32
	var phKeyHandle C.SGD_HANDLE
	err = C.SDFGenerateAgreemetDataAndKeyWithSM9(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyLen),hid,&pucResponseID,C.SGD_UINT32(uiResponseIDLen),&pucSponsorID,C.SGD_UINT32(uiSponsorIDLen),C.SGD_UINT32(uiKeyIndex),&pucResponsePrivateKey,&pucPublicKey,&pucSponsorTmpPublicKey,&pucResponseTmpPublicKey,&pucHashSB,&puiSBLen,&pucHashS2,&puiS2Len,&phKeyHandle)
	privateKey := ConvertToSM9refEncUserPrivateKeyGo(pucResponsePrivateKey)
	publicKey := ConvertToSM9refEncMasterPublicKeyGo(pucPublicKey)
	sponsorTmpPublicKey := ConvertToSM9refEncMasterPublicKeyGo(pucSponsorTmpPublicKey)
	responseTmpPublicKey := ConvertToSM9refEncMasterPublicKeyGo(pucResponseTmpPublicKey)
	responseID:= C.GoBytes(unsafe.Pointer(pucResponseID), C.int(uiResponseIDLen))
	C.free(unsafe.Pointer(pucResponseID))
	sponsorID:= C.GoBytes(unsafe.Pointer(pucSponsorID), C.int(uiSponsorIDLen))
	C.free(unsafe.Pointer(pucSponsorID))
	hashSB:= C.GoBytes(unsafe.Pointer(pucHashSB), C.int(uiResponseIDLen))
	C.free(unsafe.Pointer(pucHashSB))
	hashS2:= C.GoBytes(unsafe.Pointer(pucHashS2), C.int(uiSponsorIDLen))
	C.free(unsafe.Pointer(pucHashS2))
	return responseID,sponsorID,privateKey,publicKey,sponsorTmpPublicKey,responseTmpPublicKey,hashSB,hashS2,KeyHandleType(phKeyHandle),ToError(err)
}

func (c *Ctx)SDFGenerateKeyWithSM9(sessionHandle SessionHandleType,uiKeyLen uint,uiSponsorIDLen uint,uiResponseIDLen uint,uiKeyIndex uint,uiSBLen uint,hAgreementHandle  AgreementHandleType)([]byte,[]byte,core.SM9refEncUserPrivateKey,core.SM9refEncMasterPublicKey,core.SM9refEncMasterPublicKey,[]byte,[]byte,KeyHandleType,error){
	var err C.SGD_RV
	var hid C.SGD_UCHAR
	var pucSponsorID C.SGD_UCHAR_PRT
	var pucResponseID C.SGD_UCHAR_PRT
	var pucSponsorPrivateKey C.SM9refEncUserPrivateKey
	var pucPublicKey C.SM9refEncMasterPublicKey
	var pucResponseTmpPublicKey C.SM9refEncMasterPublicKey
	var pucHashSB C.SGD_UCHAR_PRT
	var pucHashSA C.SGD_UCHAR_PRT
	var puiSALen C.SGD_UINT32
	var phKeyHandle C.SGD_HANDLE
	err = C.SDFGenerateKeyWithSM9(c.libHandle,C.SGD_HANDLE(sessionHandle),C.SGD_UINT32(uiKeyLen),hid,&pucSponsorID,C.SGD_UINT32(uiSponsorIDLen),&pucResponseID,C.SGD_UINT32(uiResponseIDLen),C.SGD_UINT32(uiKeyIndex),&pucSponsorPrivateKey,&pucPublicKey,&pucResponseTmpPublicKey,&pucHashSB,C.SGD_UINT32(uiSBLen),&pucHashSA,&puiSALen,C.SGD_HANDLE(hAgreementHandle),&phKeyHandle)
	privateKey:=ConvertToSM9refEncUserPrivateKeyGo(pucSponsorPrivateKey)
	publicKey := ConvertToSM9refEncMasterPublicKeyGo(pucPublicKey)
	responseTmpPublicKey := ConvertToSM9refEncMasterPublicKeyGo(pucResponseTmpPublicKey)
	sponsorID:= C.GoBytes(unsafe.Pointer(pucSponsorID), C.int(uiSponsorIDLen))
	C.free(unsafe.Pointer(pucSponsorID))
	responseID:= C.GoBytes(unsafe.Pointer(pucResponseID), C.int(uiResponseIDLen))
	C.free(unsafe.Pointer(pucResponseID))
	hashSB:= C.GoBytes(unsafe.Pointer(pucHashSB), C.int(uiSBLen))
	C.free(unsafe.Pointer(pucHashSB))
	hashSA:= C.GoBytes(unsafe.Pointer(pucHashSA), C.int(puiSALen))
	C.free(unsafe.Pointer(pucHashSA))
	return sponsorID,responseID,privateKey,publicKey,responseTmpPublicKey,hashSB,hashSA,KeyHandleType(phKeyHandle),ToError(err)
}

func (c *Ctx)SDFGenerateKeyVerifySM9(sessionHandle SessionHandleType,uiS2Len uint,uiSALen uint)([]byte,[]byte,error){
	var err C.SGD_RV
	var pHashS2 C.SGD_UCHAR_PRT
	var pHashSA C.SGD_UCHAR_PRT
	err = C.SDFGenerateKeyVerifySM9(c.libHandle,C.SGD_HANDLE(sessionHandle),&pHashS2,C.SGD_UINT32(uiS2Len),&pHashSA,C.SGD_UINT32(uiSALen))
	hashS2:= C.GoBytes(unsafe.Pointer(pHashS2), C.int(uiS2Len))
	C.free(unsafe.Pointer(pHashS2))
	hashSA:= C.GoBytes(unsafe.Pointer(pHashSA), C.int(uiSALen))
	C.free(unsafe.Pointer(pHashSA))
	return hashS2,hashSA,ToError(err)
}