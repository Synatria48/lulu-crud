$(document).ready(function() {

    PopulateProvinces();

    $('#submit').click(function() {

        let pernyataanKebenaran = $('#pernyataanKebenaran').prop('checked');
        let formData = new FormData($('#multiple_section_form')[0])
        
        if (!pernyataanKebenaran) {
        alert("MOHON CEKLIS DULU PERNYATAANNYA")
        } else {
            console.log(formData);

            $.ajax({
                url: '/users-create',
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

        if ($('#provinsi :selected').attr("data_id") == "") {
            select.empty();
            select.append($('<option>', {
                value: '',
                text: '-- Pilih Kabupaten/Kota --'
            }).attr('data_id', ''));

            var select2 = $('#kecamatan');
            select2.empty();
            select2.append($('<option>', {
                value: '',
                text: '-- Pilih Kecamatan --'
            }).attr('data_id', ''));

            return;
        }

        $.getJSON(`https://www.emsifa.com/api-wilayah-indonesia/api/regencies/${$('#provinsi :selected').attr("data_id")}.json`, function (data) {
            select.append($('<option>', {
                value: '',
                text: '-- Pilih Kabupaten/Kota --'
            }).attr('data_id', ''));

            $.each(data, function (index, regencies) {
                select.append($('<option>', {
                    value: regencies.name,
                    text: regencies.name
                }).attr('data_id', regencies.id));
            });
        });
    });

    $("#kabupaten").on( "change", function() {
        var select = $('#kecamatan');
        select.empty();

        if ($('#kabupaten :selected').attr("data_id") == "") {
            select.empty();
            select.append($('<option>', {
                value: '',
                text: '-- Pilih Kecamatan --'
            }).attr('data_id', ''));

            return;
        }

        $.getJSON(`https://www.emsifa.com/api-wilayah-indonesia/api/districts/${$('#kabupaten :selected').attr("data_id")}.json`, function (data) {
            select.append($('<option>', {
                value: '',
                text: '-- Pilih Kecamatan --'
            }).attr('data_id', ''));

            $.each(data, function (index, district) {
                select.append($('<option>', {
                    value: district.name,
                    text: district.name
                }));
            });
        });
    });
});