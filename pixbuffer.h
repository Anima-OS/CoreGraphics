/*
 * Copyright (c) 2018, Mark "Happy-Ferret" Bauermeister
 *
 * This software may be modified and distributed under the terms
 * of the BSD license.  See the LICENSE file for details.
*/

static gboolean
_gdk_pixbuf_save(GdkPixbuf *pixbuf,
const char *filename, const char *format, GError ** err, const char *quality)
{
	return gdk_pixbuf_save(pixbuf, filename, format, err, "quality", quality, NULL);
}
