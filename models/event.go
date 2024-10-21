/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   event.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/21 21:31:26 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/21 21:40:42 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package models

import "time"

type Event struct {
	ID int
	Name string `binding: required`
	Description string `binding: required`
	Location string `binding: required`
	DateTime time.Time `binding: required`
	UserID int
}

var events = []Event{}

func (e *Event) New() {
	
}
func (e *Event) Save() {
	events = append(events, *e)
}

func GetAllEvents() []Event {
	return events;
}