package entity

type Feedback struct {
	ID                                int    `gorm:"primaryKey;autoIncrement" json:"id"`
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

func (Feedback) TableName() string {
	return "feedback" // Nome da tabela como no banco de dados
}
