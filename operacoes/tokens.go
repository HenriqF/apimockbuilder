package ops

import (
	"errors"
	"fmt"
	"unicode"
)

// ----
type token_tipo int

const (
	token_nome token_tipo = iota
	token_numero
	token_op
	token_delim
	token_default
	token_ponto
	token_string
)

type token struct {
	valor string
	tipo  token_tipo
}

var token_index = 0
var string_token = ""
var len_string_token = 0

//---

func is_token_operator(r rune) bool {
	if r == '+' || r == '-' || r == '*' || r == '/' {
		return true
	}
	if r == '=' || r == '>' || r == '<' || r == '!' {
		return true
	}
	return false
}
func get_token_tipo(input rune) token_tipo {
	if input == '.' {
		return token_ponto
	}

	if unicode.IsLetter(input) || input == '_' {
		return token_nome
	}

	if unicode.IsDigit(input) {
		return token_numero
	}

	if is_token_operator(input) {
		return token_op
	}

	if input == '(' || input == ')' || input == '\'' || input == '"' {
		return token_delim
	}

	return token_default
}

// para ler os tokens de tipo numero ou nome. considera o ponto como parte do token
func ler_continuo(tipo token_tipo) (token, error) {
	valor := ""
	for ; token_index < len_string_token; token_index++ {
		token_tipo := get_token_tipo(rune(string_token[token_index]))
		if token_tipo != tipo && token_tipo != token_ponto {
			break
		}
		valor += string(string_token[token_index])
	}
	token_index--
	return token{valor, tipo}, nil
}

func ler_string(delim rune) (token, error) {
	valor := ""
	token_index++
	for ; token_index < len_string_token; token_index++ {
		if rune(string_token[token_index]) == delim {
			break
		}

		valor += string(string_token[token_index])
	}
	if rune(string_token[token_index]) != delim {
		return token{}, errors.New("string nao terminada")
	}
	return token{valor, token_string}, nil
}

// se conseguir, le o proximo porraloca e incrementa o contador
func ler_proximo() rune {
	if token_index+1 < len_string_token {
		token_index++
		return rune(string_token[token_index])
	}
	return 0
}

func get_tokens(input string) ([]token, error) {
	string_token = input
	token_index = 0
	len_string_token = len(string_token)

	var tokens []token

	for ; token_index < len_string_token; token_index++ {
		char := rune(string_token[token_index])
		token_tipo := get_token_tipo(char)

		switch token_tipo {
		case token_numero, token_nome:
			novo, err := ler_continuo(token_tipo)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, novo)

		default:
			switch char {
			case '.':
				tokens = append(tokens, token{".", token_ponto})
			case '-':
				resultado := string(char)

				if len(tokens) == 0 {
					resultado = "u-"
				} else {
					ultimo := tokens[len(tokens)-1].tipo
					if ultimo == token_op || ultimo == token_nome {
						resultado = "u-"
					}
				}
				tokens = append(tokens, token{resultado, token_op})

			case '+', '*', '/':
				tokens = append(tokens, token{string(char), token_op})
			case '=', '!', '>', '<':
				if ler_proximo() == '=' {
					tokens = append(tokens, token{string(char) + "=", token_op})
				} else {
					tokens = append(tokens, token{string(char), token_op})
				}
			case '\'', '"':
				novo, err := ler_string(char)
				if err != nil {
					return nil, err
				}
				tokens = append(tokens, novo)
			case '(', ')':
				tokens = append(tokens, token{string(char), token_delim})
			}

		}

	}

	fmt.Println("sem ordem:", tokens)
	return tokens, nil
}

func ordenar_tokens(input []token) ([]token, error) {
	var stack []token
	var final []token

	for _, t := range input {
		if t.tipo == token_nome || t.tipo == token_numero || t.tipo == token_string {
			final = append(final, t)
			continue
		}

		if len(stack) <= 0 {
			stack = append(stack, t)
			continue
		}

		if t.valor == ")" {
			for {
				if len(stack) == 0 {
					return final, errors.New("parenteses desbalanceados")
				}

				if stack[len(stack)-1].valor == "(" {
					stack = stack[:len(stack)-1]
					break
				}
				final = append(final, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			continue
		}

		if t.valor == "(" {
			stack = append(stack, t)
			continue
		}

		for {
			p_self := get_token_precedence(t)
			p_past := get_token_precedence(stack[len(stack)-1])

			if p_self <= p_past {
				final = append(final, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			if len(stack) == 0 || p_self > p_past {
				stack = append(stack, t)
				break
			}
		}

	}

	for {
		if len(stack) == 0 {
			break
		}

		if stack[len(stack)-1].valor == "(" {
			return final, errors.New("parenteses desbalanceados")
		}

		final = append(final, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	fmt.Println("com ordem:", final)
	return final, nil
}
