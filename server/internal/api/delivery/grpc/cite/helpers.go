package cite

import "test_server/proto/cite"

// handleRequestData handles request data.
// returns true if the request is done, false if the request is not done.
func (h Handler) handleRequestData(stream cite.CiteService_GetCiteServer, data request, nonce string) (bool, error) {
	err := data.err
	if err != nil {
		if err.Error() == eof {
			return true, nil
		}
		return false, err
	}

	// Process a request
	req := data.req
	switch {
	case h.ddosEnabled && req.Pow == "":
		// If the POW is empty, ask to solve the POW
		if err = stream.Send(&cite.CiteResponse{PowRiddle: h.powService.GetPow(nonce)}); err != nil {
			h.log.Error("send pow", err)
			return false, err
		}
	default:
		if h.ddosEnabled && !h.powService.ValidatePOW(h.powService.GetPow(nonce), req.Pow) {
			return false, ErrorInvalidPOW
		}

		// If the POW is valid, get the Cite and send it to the client
		var citePayload string
		citePayload, err = h.citeService.GetCite()
		if err != nil {
			h.log.Error("get cite", err)
			return false, err
		}

		if err = stream.Send(&cite.CiteResponse{Cite: citePayload}); err != nil {
			h.log.Error("send cite", err)
			return false, err
		}

		return true, nil
	}

	return false, nil
}
