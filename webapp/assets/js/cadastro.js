$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario(evento) {
    evento.preventDefault();

    if ($('#senha').val() != $('#confirmar-senha').val()) {
        alert("As senhas não são iguais");
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val(),
        }
    }).done(function() {
        alert("Usuário cadastrado com sucesso!")
        window.location = "/login"
    }).fail(function(erro) {
        console.log(erro)
        alert("Erro ao cadastrar o usuário!")
    });
}