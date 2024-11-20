//go:build darwin && !ios
// +build darwin,!ios

#import <Foundation/Foundation.h>

char* preferredLocalization() {
	NSString *localization = [NSBundle mainBundle].preferredLocalizations.firstObject;

	return (char *) [localization UTF8String];
}

int preferredLocalizations(char ***list) {
	NSArray *locs = [NSBundle mainBundle].preferredLocalizations;
	int nlocs = [locs count];
	char **r = calloc(nlocs, sizeof(char *));

	if (r == NULL)
		return -1;

	for (int n = 0; n < nlocs; n++) {
		r[n] = strdup([locs[n] UTF8String]);
		if (r[n] == NULL)
			goto fail;
	}

	*list = r;

	return nlocs;

fail:
	for (int n = 0; n < nlocs; n++) {
		if (r[n] != NULL)
			free(r[n]);
	}
	free(r);
	return -1;
}
