function PopulateProvinces() {
    $.getJSON('https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json', function (data) {
        var select = $('#provinsi');
        select.empty();

        select.append($('<option>', {
            value: '',
            text: '-- Pilih Provinsi --'
        }).attr('data_id', ''));

        $.each(data, function (index, province) {
        select.append($('<option>', {
            value: province.name,
            text: province.name
        }).attr('data_id', province.id));
        });
    });
}