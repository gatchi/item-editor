/* PSO ItemPMT Browser Editor
 * Copyright 2017 Christen Gottschlich
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package itempmt

// Offset from end of ItemPMT.bin file to main pointer table.
const START int64 = -16

// Offsets for main table, in bytes.
var Offset = map[string]int64 {
	"weapon": 0,
	"armor": 4,
	"units": 8,
	"item": 12,
	"mag": 16,
	"AA": 20,
	"photon-color": 24,
	"weapon-range": 28,
	"weapon-sales-div": 32,
	"sales-div": 36,
	"mag-feed": 40,
	"star-value": 44,
	"special-data": 48,
	"weapon-SFX": 52,
	"stat-boost": 56,
	"shield-sfx": 60,
	"tech-use": 64,
	"combination": 68,
	// "unknown-table": 72,
	"tech-bosts": 76,
	"unwrap": 80,
	"unsealable": 84,
	"ranged-special": 88,
}
