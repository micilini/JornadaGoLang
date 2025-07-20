package arch

import (
    "os"
    "path/filepath"
    "runtime"
    "testing"
)

// mapeia cada SO ao path que deve existir
var requiredPaths = map[string]string{
    "windows": `C:\Windows`,
    "linux":   "/usr",
    "darwin":  "/Applications",
}

func TestArchSupport(t *testing.T) {
    osName := runtime.GOOS   // "windows", "linux" ou "darwin"
    arch := runtime.GOARCH   // "amd64", "386", "arm64", etc.

    t.Logf("SO detectado: %s  |  ARCH detectada: %s", osName, arch)

    // 1) Verifica se suportamos esse SO
    basePath, supported := requiredPaths[osName]
    if !supported {
        t.Skipf("SO não suportado: %s — este programa não roda aqui", osName)
    }

    // 2) Monta um exemplo de subpasta dentro do path base
    //    (por exemplo, "%SystemRoot%\System32" no Windows ou "/usr/local/bin" no Unix)
    var testPath string
    switch osName {
    case "windows":
        testPath = filepath.Join(basePath, "System32")
    case "linux", "darwin":
        testPath = filepath.Join(basePath, "local", "bin")
    }

    // 3) Verifica se a pasta existe
    info, err := os.Stat(testPath)
    if err != nil {
        if os.IsNotExist(err) {
            t.Errorf("Pasta esperada não encontrada em %s para %s/%s", testPath, osName, arch)
        } else {
            t.Errorf("Erro checando %s: %v", testPath, err)
        }
        return
    }
    if !info.IsDir() {
        t.Errorf("Caminho existe mas não é pasta: %s", testPath)
        return
    }

    // 4) Se chegou até aqui, está tudo OK
    t.Logf("✔ Pasta %s encontrada. Arquitetura %s/%s suportada.", testPath, osName, arch)
}