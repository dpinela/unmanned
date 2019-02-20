package main

var defaultStylesheet = []byte(
	`/* $Id: mandoc.css,v 1.42 2018/12/04 06:11:49 schwarze Exp $ */
/*
 * Standard style sheet for mandoc(1) -Thtml and man.cgi(8).
 *
 * Written by Ingo Schwarze <schwarze@openbsd.org>.
 * I place this file into the public domain.
 * Permission to use, copy, modify, and distribute it for any purpose
 * with or without fee is hereby granted, without any conditions.
 */

/* Global defaults. */

html {		max-width: 65em; }
body {		font-family: Helvetica,Arial,sans-serif; }
table {		margin-top: 0em;
		margin-bottom: 0em;
		border-collapse: collapse; }
/* Some browsers set border-color in a browser style for tbody,
 * but not for table, resulting in inconsistent border styling. */
tbody {		border-color: inherit; }
tr {		border-color: inherit; }
td {		vertical-align: top;
		padding-left: 0.2em;
		padding-right: 0.2em;
		border-color: inherit; }
ul, ol, dl {	margin-top: 0em;
		margin-bottom: 0em; }
li, dt {	margin-top: 1em; }

.permalink {	border-bottom: thin dotted;
		color: inherit;
		font: inherit;
		text-decoration: inherit; }
* {		clear: both }

/* Search form and search results. */

fieldset {	border: thin solid silver;
		border-radius: 1em;
		text-align: center; }
input[name=expr] {
		width: 25%; }

table.results {	margin-top: 1em;
		margin-left: 2em;
		font-size: smaller; }

/* Header and footer lines. */

table.head {	width: 100%;
		border-bottom: 1px dotted #808080;
		margin-bottom: 1em;
		font-size: smaller; }
td.head-vol {	text-align: center; }
td.head-rtitle {
		text-align: right; }

table.foot {	width: 100%;
		border-top: 1px dotted #808080;
		margin-top: 1em;
		font-size: smaller; }
td.foot-os {	text-align: right; }

/* Sections and paragraphs. */

.manual-text {
		margin-left: 3.8em; }
.Nd {		display: inline; }
.Sh {		margin-top: 1.2em;
		margin-bottom: 0.6em;
		margin-left: -3.2em;
		font-size: 110%; }
.Ss {		margin-top: 1.2em;
		margin-bottom: 0.6em;
		margin-left: -1.2em;
		font-size: 105%; }
.Pp {		margin: 0.6em 0em; }
.Sx { }
.Xr { }

/* Displays and lists. */

.Bd { }
.Bd-indent {	margin-left: 3.8em; }

.Bl-bullet {	list-style-type: disc;
		padding-left: 1em; }
.Bl-bullet > li { }
.Bl-dash {	list-style-type: none;
		padding-left: 0em; }
.Bl-dash > li:before {
		content: "\2014  "; }
.Bl-item {	list-style-type: none;
		padding-left: 0em; }
.Bl-item > li { }
.Bl-compact > li {
		margin-top: 0em; }

.Bl-enum {	padding-left: 2em; }
.Bl-enum > li { }
.Bl-compact > li {
		margin-top: 0em; }

.Bl-diag { }
.Bl-diag > dt {
		font-style: normal;
		font-weight: bold; }
.Bl-diag > dd {
		margin-left: 0em; }
.Bl-hang { }
.Bl-hang > dt { }
.Bl-hang > dd {
		margin-left: 5.5em; }
.Bl-inset { }
.Bl-inset > dt { }
.Bl-inset > dd {
		margin-left: 0em; }
.Bl-ohang { }
.Bl-ohang > dt { }
.Bl-ohang > dd {
		margin-left: 0em; }
.Bl-tag {	margin-left: 5.5em; }
.Bl-tag > dt {
		float: left;
		margin-top: 0em;
		margin-left: -5.5em;
		padding-right: 0.5em;
		vertical-align: top; }
.Bl-tag > dd {
		clear: right;
		width: 100%;
		margin-top: 0em;
		margin-left: 0em;
		vertical-align: top;
		overflow: auto; }
.Bl-compact > dt {
		margin-top: 0em; }

.Bl-column { }
.Bl-column > tbody > tr { }
.Bl-column > tbody > tr > td {
		margin-top: 1em; }
.Bl-compact > tbody > tr > td {
		margin-top: 0em; }

.Rs {		font-style: normal;
		font-weight: normal; }
.RsA { }
.RsB {		font-style: italic;
		font-weight: normal; }
.RsC { }
.RsD { }
.RsI {		font-style: italic;
		font-weight: normal; }
.RsJ {		font-style: italic;
		font-weight: normal; }
.RsN { }
.RsO { }
.RsP { }
.RsQ { }
.RsR { }
.RsT {		text-decoration: underline; }
.RsU { }
.RsV { }

.eqn { }
.tbl td {	vertical-align: middle; }

.HP {		margin-left: 3.8em;
		text-indent: -3.8em; }

/* Semantic markup for command line utilities. */

table.Nm { }
code.Nm {	font-style: normal;
		font-weight: bold;
		font-family: inherit; }
.Fl {		font-style: normal;
		font-weight: bold;
		font-family: inherit; }
.Cm {		font-style: normal;
		font-weight: bold;
		font-family: inherit; }
.Ar {		font-style: italic;
		font-weight: normal; }
.Op {		display: inline; }
.Ic {		font-style: normal;
		font-weight: bold;
		font-family: inherit; }
.Ev {		font-style: normal;
		font-weight: normal;
		font-family: monospace; }
.Pa {		font-style: italic;
		font-weight: normal; }

/* Semantic markup for function libraries. */

.Lb { }
code.In {	font-style: normal;
		font-weight: bold;
		font-family: inherit; }
a.In { }
.Fd {		font-style: normal;
		font-weight: bold;
		font-family: inherit; }
.Ft {		font-style: italic;
		font-weight: normal; }
.Fn {		font-style: normal;
		font-weight: bold;
		font-family: inherit; }
.Fa {		font-style: italic;
		font-weight: normal; }
.Vt {		font-style: italic;
		font-weight: normal; }
.Va {		font-style: italic;
		font-weight: normal; }
.Dv {		font-style: normal;
		font-weight: normal;
		font-family: monospace; }
.Er {		font-style: normal;
		font-weight: normal;
		font-family: monospace; }

/* Various semantic markup. */

.An { }
.Lk { }
.Mt { }
.Cd {		font-style: normal;
		font-weight: bold;
		font-family: inherit; }
.Ad {		font-style: italic;
		font-weight: normal; }
.Ms {		font-style: normal;
		font-weight: bold; }
.St { }
.Ux { }

/* Physical markup. */

.Bf {		display: inline; }
.No {		font-style: normal;
		font-weight: normal; }
.Em {		font-style: italic;
		font-weight: normal; }
.Sy {		font-style: normal;
		font-weight: bold; }
.Li {		font-style: normal;
		font-weight: normal;
		font-family: monospace; }

/* Overrides to avoid excessive margins on small devices. */

@media (max-width: 37.5em) {
.manual-text {
		margin-left: 0.5em; }
.Sh, .Ss {	margin-left: 0em; }
.Bd-indent {	margin-left: 2em; }
.Bl-hang > dd {
		margin-left: 2em; }
.Bl-tag {	margin-left: 2em; }
.Bl-tag > dt {
		margin-left: -2em; }
.HP {		margin-left: 2em;
		text-indent: -2em; }
}

`)
