$(document).ready(function() {

    PopulateProvinces();

    $('#submit').click(function() {
        var formData = {};

        formData.nama = $('#nama').val();
        formData.nik = $('#nik').val();
        formData.tanggalLahir = $('#tanggalLahir').val();
        formData.alamat = $('#alamat').val();

        formData.provinsi = $('#provinsi :selected').text();
        formData.kabupaten = $('#kabupaten :selected').text();
        formData.kecamatan = $('#kecamatan :selected').text();

        formData.nomorHp = $('#nomorHp').val();
        formData.email = $('#email').val();

        formData.jenisKelamin = $('input[name="jenisKelamin"]:checked').val();

        formData.tingkatPendidikan = $('#tingkatPendidikan').val();

        formData.uploadFoto = $('#uploadFoto')[0].files[0];
        formData.uploadDokumen = $('#uploadDokumen')[0].files;

        formData.pernyataanKebenaran = $('#pernyataanKebenaran').prop('checked');

        if (!formData.pernyataanKebenaran) {
        alert("MOHON CEKLIS DULU PERNYATAANNYA")
        } else {
            $.ajax({
                url: '/users',
                type: 'POST',
                data: formData,
                contentType: false,
                processData: false,
                success: function (response) {
                    console.log(response);
                },
                error: function (error) {
                    console.error(error);
                }
            });
        }
    });

    $('#reset').click(function() {
        $('#nama').val('');
        $('#nik').val('');
        $('#tanggalLahir').val('');
        $('#alamat').val('');
        $('#provinsi').val('');
        $('#kabupaten').val('');
        $('#kecamatan').val('');
        $('#nomorHp').val('');
        $('#email').val('');
        $('input[name="jenisKelamin"]').prop('checked', false);
        $('#tingkatPendidikan').val('');
        $('#uploadFoto').val('');
        $('#uploadDokumen').val('');
        $('#pernyataanKebenaran').prop('checked', false);
    });

    $("#provinsi").on( "change", function() {
        var select = $('#kabupaten');
        select.empty();

        if ($('#provinsi :selected').val() == "") {
            select.empty();
            select.append($('<option>', {
                value: '',
                text: '-- Pilih Kabupaten/Kota --'
            }));

            var select2 = $('#kecamatan');
            select2.empty();
            select2.append($('<option>', {
                value: '',
                text: '-- Pilih Kecamatan --'
            }));

            return;
        }

        $.getJSON(`https://www.emsifa.com/api-wilayah-indonesia/api/regencies/${$('#provinsi :selected').val()}.json`, function (data) {
            select.append($('<option>', {
                value: '',
                text: '-- Pilih Kabupaten/Kota --'
            }));

            $.each(data, function (index, regencies) {
                select.append($('<option>', {
                    value: regencies.id,
                    text: regencies.name
                }));
            });
        });
    });

    $("#kabupaten").on( "change", function() {
        var select = $('#kecamatan');
        select.empty();

        if ($('#kabupaten :selected').val() == "") {
            select.empty();
            select.append($('<option>', {
                value: '',
                text: '-- Pilih Kecamatan --'
            }));

            return;
        }

        $.getJSON(`https://www.emsifa.com/api-wilayah-indonesia/api/districts/${$('#kabupaten :selected').val()}.json`, function (data) {
            select.append($('<option>', {
                value: '',
                text: '-- Pilih Kecamatan --'
            }));

            $.each(data, function (index, district) {
                select.append($('<option>', {
                    value: district.id,
                    text: district.name
                }));
            });
        });
    });
});