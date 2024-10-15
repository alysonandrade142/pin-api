package dto

type FeedbackDTO struct {
	FuncionarioID                     int    `json:"funcionario_id"`
	Feedback                          string `json:"feedback"`
	ComentariosFeedback               string `json:"comentarios_feedback"`
	InteracaoGestor                   string `json:"interacao_gestor"`
	ComentariosInteracaoGestor        string `json:"comentarios_interacao_gestor"`
	ClarezaCarreira                   string `json:"clareza_carreira"`
	ComentariosClarezaCarreira        string `json:"comentarios_clareza_carreira"`
	ExpectativaPermanencia            string `json:"expectativa_permanencia"`
	ComentariosExpectativaPermanencia string `json:"comentarios_expectativa_permanencia"`
	ENPS                              int    `json:"enps"`
	AbertaENPS                        string `json:"aberta_enps"`
}
