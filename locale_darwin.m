//go:build darwin && !ios
// +build darwin,!ios

#import <Foundation/Foundation.h>

char* preferredLocalization() {
	NSString *localization = [NSBundle mainBundle].preferredLocalizations.firstObject;

	return (char *) [localization UTF8String];
}
