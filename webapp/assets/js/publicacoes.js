$('#nova-publicacao').on('submit', criarPublicacao);

$(document).on('click', '.curtir-publicacao', curtirPublicacao)
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao)

$('#atualizar-publicacao').on('click', atualizarPublicacao);

function criarPublicacao(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function() {
        alert("Erro ao criar a publicacao");
    })
} 

function curtirPublicacao(evento) {
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

    elementoClicado.prop('disabled', true);

    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST",
    }).done(function() {
        const contadorDeCurtidas = elementoClicado.next('span')
        const quatidadeDeCurtidas = parseInt(contadorDeCurtidas.text())

        contadorDeCurtidas.text(quatidadeDeCurtidas + 1)

        elementoClicado.addClass('descurtir-publicacao')
        elementoClicado.addClass('text-danger')
        elementoClicado.removeClass('curtir-publicacao')


    }).fail(function() {
        alert("Erro ao curtir publicacao")
    }).always(function() {
        elementoClicado.prop('disabled', false);
    })
}

function descurtirPublicacao(evento) {
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

    elementoClicado.prop('disabled', true);

    $.ajax({
        url: `/publicacoes/${publicacaoId}/descurtir`,
        method: "POST",
    }).done(function() {
        const contadorDeCurtidas = elementoClicado.next('span')
        const quatidadeDeCurtidas = parseInt(contadorDeCurtidas.text())

        contadorDeCurtidas.text(quatidadeDeCurtidas - 1)

        elementoClicado.addClass('curtir-publicacao')
        elementoClicado.removeClass('text-danger')
        elementoClicado.removeClass('descurtir-publicacao')


    }).fail(function() {
        alert("Erro ao descurtir publicacao")
    }).always(function() {
        elementoClicado.prop('disabled', false);
    })
}

function atualizarPublicacao() {
    $(this).prop('disabled', true)

    const publicacaoId = $(this).data('publicacao-id')
    
    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "PUT",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function() {
        alert("Publicação editada com sucesso!");
        window.location = "/home";
    }).fail(function() {
        alert("Erro ao editar a publicação!")
    }).always(function() {
        $('#atualizar-publicacao').prop('disabled', false)
    })
}