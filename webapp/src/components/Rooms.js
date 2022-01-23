export function Room(school, major, room_num, room_desc) {
	this.school = school;
	this.major = major;
	this.room_num = room_num;
	this.room_desc = room_desc;
}

export function RoomList(CurrList, PrevList, listDesc) {
	this.CurrList = CurrList;
	this.listDesc = listDesc;
	// If prevlist is not null
	this.PrevList = PrevList
		? new RoomList(PrevList.CurrList, PrevList.PrevList, PrevList.listDesc)
		: null;
}

export const Schools = new RoomList(
	[
		new Room("_GLOB", null, "0", "Global Chatroom"),
		new Room("BAR01", null, null, "Baruch College"),
		new Room("BCC01", null, null, "Bronx Community College"),
		new Room("BKL01", null, null, "Brooklyn College"),
		new Room("BMC01", null, null, "Borough of Manhattan Community College"),
		new Room("CSI01", null, null, "College of Staten Island"),
		new Room("CTY01", null, null, "The City College of New York"),
		new Room("GRD01", null, null, "CUNY Graduate Center"),
		new Room("HOS01", null, null, "Hostos Community College"),
		new Room("HTR01", null, null, "Hunter College"),
		new Room("JJC01", null, null, "John Jay College of Criminal Justice"),
		new Room("KCC01", null, null, "Kingsborough Community College"),
		new Room("LAG01", null, null, "LaGuardia Community College"),
		new Room("LAW01", null, null, "CUNY School of Law"),
		new Room("LEH01", null, null, "Lehman College"),
		new Room("MEC01", null, null, "Macaulay Honors College"),
		new Room("MED01", null, null, "Medgar Evans College"),
		new Room("NCC01", null, null, "Guttman Community College"), //May be wrong
		new Room("NYT01", null, null, "New York City College of Technology"), //May be wrong
		new Room("QCC01", null, null, "Queensborough Community College"),
		new Room("QNS01", null, null, "Queens College"),
		new Room("SPS01", null, null, "CUNY School of Professional Studies"),
		new Room("YRK01", null, null, "York College"),
		new Room("SLU01", null, null, "CUNY School of Labor and Urban Studies"),
		new Room(
			"SOJ01",
			null,
			null,
			"Craig Newmark Graduate School of Journalism"
		),
		new Room(
			"SPH01",
			null,
			null,
			"CUNY Graduate School of Public Health and Health Policy"
		),
	],
	null,
	"Schools"
);
