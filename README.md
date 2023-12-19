# Web Service in Go

Bu proje, basit bir Go dilindeki web servisi örneğidir. Proje, HTTP isteklerini işlemek, JWT (JSON Web Token) kullanarak kimlik doğrulama sağlamak, PostgreSQL veritabanında görev oluşturmak ve Prometheus ile metrik toplamak gibi temel özellikleri içermektedir.

## Kurulum

1. **Gereksinimler:**
    - Go dilini [buradan](https://golang.org/dl/) indirip yükleyin.
    - PostgreSQL veritabanını yükleyin ve bağlantı bilgilerini `database.go` dosyasında güncelleyin.
    - Gerekli Go paketlerini yüklemek için terminal veya komut istemcisine şu komutu yazın:

        ```bash
        go get -u ./...
        ```

2. **Çalıştırma:**
    - Proje dizininde aşağıdaki komutu çalıştırarak web servisini başlatın:

        ```bash
        go run main.go
        ```

3. **Swagger Dokümantasyonu:**
    - `swagger.json` dosyasını kullanarak API'ye dair Swagger dokümantasyonunu görmek için [Swagger Editor](https://editor.swagger.io/) veya benzer bir aracı kullanabilirsiniz.

## Kullanım

- Ana sayfa için GET isteği yapmak için:

    ```bash
    curl http://localhost:8080/
    ```

- POST isteği yapmak için:

    ```bash
    curl -X POST -d "username=testuser&password=testpassword" http://localhost:8080/post
    ```

- Token almak için:

    ```bash
    curl http://localhost:8080/token
    ```

## Katkıda Bulunma

1. Bu depoyu fork edin.
2. Yeni bir özellik eklemek veya bir hata düzeltmek için dal (branch) oluşturun.
3. Yaptığınız değişiklikleri commit'leyin.
4. Fork edilmiş repoya dalınızı (branch) push edin.
5. Pull isteği (pull request) oluşturun.

## Lisans

Bu proje MIT lisansı altında lisanslanmıştır - [LICENSE.md](LICENSE.md) dosyasına bakın.
